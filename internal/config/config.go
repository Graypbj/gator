package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (*Config, error) {
	filePath, err := getFilePath()
	if err != nil {
		return &Config{}, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return &Config{}, err
	}
	defer file.Close()

	var newConfig Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&newConfig); err != nil {
		return &Config{}, err
	}

	return &newConfig, nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName

	filePath, err := getFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonConfig, err := json.Marshal(c)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonConfig)

	return nil
}

func getFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filePath := homeDir + "/" + configFileName
	return filePath, nil
}
