package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func Load(file string) Config {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err.Error())
	}
	var c Config
	err = yaml.Unmarshal(raw, &c)
	if err != nil {
		panic(err.Error())
	}
	return c
}

type Config struct {
	Address string   `yaml:"address"`
	DB      DBConfig `yaml:"database"`
}

type DBConfig struct {
	Host     string `yaml:"hostname"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
