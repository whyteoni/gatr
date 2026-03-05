package config

import (
	"encoding/json"
	"os"

	"github.com/google/uuid"
)

const configFileName = ".gatrconfig.json"

type Config struct {
	DB_url 	string	`json:"db_url"`
	CurrentUser	struct{
		Name	string		`json:"name"`
		ID		uuid.UUID	`json:"id"`
	}	`json:"current_user"`
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

func (c *Config) SetUser(name string, id uuid.UUID) (err error) {
	c.CurrentUser.Name = name
	c.CurrentUser.ID = id

	file, err := os.Create(getConfigPath())
	if err != nil { return }
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(c)
}
