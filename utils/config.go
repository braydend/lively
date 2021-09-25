package utils

import (
	_ "embed"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Start    string   `yaml:"start"`
	End      string   `yaml:"end"`
	Interval int      `yaml:"interval"`
	Messages []string `yaml:"messages"`
}

//go:embed defaultConfig.yml
var defaultConfig []byte

func ParseConfigFromFile(filename string) Config {
	if filename == "" {
		return ParseConfig(defaultConfig)
	}

	return ParseConfig(ReadFile(filename))
}

func ParseConfig(configData []byte) (config Config) {
	err := yaml.Unmarshal(configData, &config)

	if err != nil {
		panic(err)
	}

	return config
}
