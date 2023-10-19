package data

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"fyne.io/fyne/v2/widget"
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
	Name, Host, Port, User, Password, DbName string
	Connection                               *sql.DB
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
	port.SetText("3306")

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
type Configs struct {
	Servers []*SerializationConnectionData
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
	resp := GetConfigs()
	for _, v := range resp.Servers {
		if v.Name == c.Name {
			return errors.New("Connection name has registered")
		}

	}
	resp.Servers = append(resp.Servers, c)
	err := os.WriteFile(configFilePath, []byte(resp.String()), os.ModePerm)
	if err != nil {
		panic(err)
	}

	return nil
}
