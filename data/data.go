package data

import (
	"encoding/json"
	"io"
	"os"

	"fyne.io/fyne/v2/widget"
)

const configFilePath = "./.config"

type ConnectionData struct {
	Name     *widget.Entry
	Host     *widget.Entry
	Port     *widget.Entry
	User     *widget.Entry
	Password *widget.Entry
}

type serializationConnectionData struct {
	Name, Host, Port, User, Password string
}

func NewConnectionData() *ConnectionData {
	return &ConnectionData{
		Name:     widget.NewEntry(),
		Host:     widget.NewEntry(),
		Port:     widget.NewEntry(),
		User:     widget.NewEntry(),
		Password: widget.NewPasswordEntry(),
	}
}

func (c *ConnectionData) Save() {
	data := c.Get()
	SaveServer(data)
}

func (c *ConnectionData) Get() *serializationConnectionData {
	return &serializationConnectionData{
		Name:     c.Name.Text,
		Host:     c.Host.Text,
		Port:     c.Port.Text,
		User:     c.User.Text,
		Password: c.Password.Text,
	}
}
func (c *ConnectionData) String() string {
	jm, _ := json.Marshal(serializationConnectionData{
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
	Servers []*serializationConnectionData
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
func SaveServer(c *serializationConnectionData) {
	resp := GetConfigs()
	resp.Servers = append(resp.Servers, c)
	err := os.WriteFile(configFilePath, []byte(resp.String()), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
