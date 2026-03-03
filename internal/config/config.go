package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatrconfig.json"

type Config struct {
	DB_url 	string	`json:"db_url"`
	CurrentUserName	string	`json:"current_user_name"`
}

func getConfigPath() (path string) {
	path, err := os.UserHomeDir()
	if err != nil { return }
	path += "/" + configFileName
	return
}

func Read() (cfg Config, err error) {
	data, err := os.ReadFile(getConfigPath())
	if err != nil { return }

	err = json.Unmarshal(data, &cfg)
	return
}

func (c *Config) SetUser(user string) (err error) {
	c.CurrentUserName = user

	file, err := os.Create(getConfigPath())
	if err != nil { return }
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(c)
}
