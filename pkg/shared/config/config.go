package config

import (
	"fmt"

	Database "anymind/pkg/shared/database"

	"github.com/spf13/viper"
)

type Config struct {
	Apps     Apps `json:"apps"`
	Database DB   `json:"database"`
}

type Apps struct {
	Name     string `json:"name"`
	HttpPort int    `json:"httpPort"`
	Version  string `json:"version"`
}

type DB struct {
	Master Database.ConfigDatabase `json:"master"`
	Slave  Database.ConfigDatabase `json:"slave"`
}

func (c *Config) AppAddress() string {
	return fmt.Sprintf(":%v", c.Apps.HttpPort)
}

func NewConfig(path string) *Config {
	fmt.Println("Try NewConfig ... ")

	viper.SetConfigFile(path)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	conf := Config{}
	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	return &conf
}
