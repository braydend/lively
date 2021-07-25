package main

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	Start    string   `yaml:"start"`
	End      string   `yaml:"end"`
	Interval int      `yaml:"interval"`
	Messages []string `yaml:"messages"`
}

func ParseConfig(configData []byte) (config Config) {
	err := yaml.Unmarshal(configData, &config)

	if err != nil {
		panic(err)
	}

	return config
}
