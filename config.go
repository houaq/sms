package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type config struct {
	ComPort    string
	BaudRate   int
	ServerHost string
	ServerPort int
}

func NewConfig(configPath string) (config, error) {
	var conf config
	_, err := toml.DecodeFile(configPath, &conf)
	if err != nil {
		return conf, fmt.Errorf("New: %s", err.Error())
	}

	return conf, nil
}
