package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetConfigFromFile(configFile string) (AppConfig, error) {
	var appConfig AppConfig
	configFileContent, err := os.ReadFile(configFile)
	if err != nil {
		return appConfig, fmt.Errorf("error while reading config file: %v", err)
	}
	err = json.Unmarshal(configFileContent, &appConfig)
	if err != nil {
		return appConfig, fmt.Errorf("error while unmashalling config file: %v", err)
	}
	return appConfig, nil
}
