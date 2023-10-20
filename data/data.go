package data

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"fyne.io/fyne/v2/widget"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Connections = map[widget.TreeNodeID]*sql.DB{}

const configFilePath = "./.config"

type ConnectionData struct {
	Name     *widget.Entry
	Host     *widget.Entry
	Port     *widget.Entry
	User     *widget.Entry
	Password *widget.Entry
}

type SerializationConnectionData struct {
	Name       string  `json:"name"`
	Host       string  `json:"host"`
	Port       string  `json:"port"`
	User       string  `json:"user"`
	Password   string  `json:"password"`
	DbName     string  `json:"db_name"`
	Connection *sql.DB `json:"connection"`
	NodeType   int     `json:"node_type"`
}

func (s *SerializationConnectionData) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", s.User, s.Password, s.Host, s.Port)
}

func NewConnectionData() *ConnectionData {
	var name = widget.NewEntry()
	var host = widget.NewEntry()
	var port = widget.NewEntry()
	var user = widget.NewEntry()
	var password = widget.NewPasswordEntry()

	//set default port
	name.SetText("local")
	host.SetText("localhost")
	user.SetText("root")
	port.SetText("3306")
	password.SetText("123456")

	name.Validator = func(s string) error {
		if s == "" {
			return errors.New("Empty Name !")
		}

		return nil
	}

	host.Validator = func(s string) error {
		if s == "" {
			return errors.New("Empty Host !")
		}

		return nil
	}

	port.Validator = func(s string) error {
		if s == "" {
			return errors.New("Empty Port !")
		}

		return nil
	}

	port.SetPlaceHolder("3306")

	user.Validator = func(s string) error {
		if s == "" {
			return errors.New("Empty User !")
		}

		return nil
	}
	return &ConnectionData{
		Name:     name,
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
	}
}

func (c *ConnectionData) Save() error {
	data := c.Get()
	return SaveServer(data)
}

func (c *ConnectionData) Get() *SerializationConnectionData {
	return &SerializationConnectionData{
		Name:     c.Name.Text,
		Host:     c.Host.Text,
		Port:     c.Port.Text,
		User:     c.User.Text,
		Password: c.Password.Text,
	}
}
func (c *ConnectionData) String() string {
	jm, _ := json.Marshal(SerializationConnectionData{
		Name:     c.Name.Text,
		Host:     c.Host.Text,
		Port:     c.Port.Text,
		User:     c.User.Text,
		Password: c.Password.Text,
	})
	return string(jm)
}

// configs

var GlobalConfigure *Configs

type Configs struct {
	Servers []*SerializationConnectionData `json:"servers"`
}

func initConfigFile() {
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		file, err := os.Create(configFilePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		_, err = file.WriteString("{}")
		if err != nil {
			return
		}
		defer file.Close()
		fmt.Println("Config File created")
	}
}

func LoadConfig() error {
	initConfigFile()
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&GlobalConfigure); err != nil {
		return err
	}

	viper.WatchConfig()

	// 配置文件更改时的回调函数
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed:", in.Name)
		if err := viper.Unmarshal(&GlobalConfigure); err != nil {
			fmt.Println("Error unmarshaling config:", err)
		}
	})

	return nil
}

func (c Configs) String() string {
	jm, _ := json.Marshal(c)
	return string(jm)
}

func GetConfigs() Configs {
	file, err := os.OpenFile(configFilePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var resp = Configs{}
	_ = json.Unmarshal(content, &resp)
	return resp
}
func SaveServer(c *SerializationConnectionData) error {
	for _, v := range GlobalConfigure.Servers {
		if v.Name == c.Name {
			return errors.New("Connection name has registered")
		}
	}
	GlobalConfigure.Servers = append(GlobalConfigure.Servers, c)
	viper.Set("servers", GlobalConfigure.Servers)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}
