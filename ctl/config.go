package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

const (
	configTOML = "config.toml"
)

type config struct {
	Username  string
	Password  string
	Cookie    string
	CSRFtoken string
}

func (c config) String() string {
	return fmt.Sprintf("Username: %s, Password: %s", c.Username, c.Password)
}

func getConfig() *config {
	cfg := new(config)
	if _, err := toml.DecodeFile(configTOML, &cfg); err != nil {
		log.Panicf(err.Error())
	}
	// log.Printf("get config: %s", cfg)
	return cfg
}
