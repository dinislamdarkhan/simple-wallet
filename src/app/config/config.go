package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configs struct {
	Cassandra CassandraConfig `json:"cassandra"`
	Logrus    Logrus          `json:"logrus"`
}

type CassandraConfig struct {
	ConnectionIP []string `json:"connection_ip"`
}

type Logrus struct {
	Level uint8 `json:"level"`
}

func GetConfigs() (*Configs, error) {
	var path string
	var configs Configs

	if os.Getenv("config") == "" {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		path = pwd + "/src/app/config/config.json"
	} else {
		path = os.Getenv("config")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configs)

	if err != nil {
		return nil, err
	}

	return &configs, nil
}
