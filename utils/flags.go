package utils

import "flag"

type flags struct {
	ConfigFile string
}

func InitFlags() flags {
	var configFile string
	flag.StringVar(&configFile, "config", "", "A YAML file to configure the app")
	flag.Parse()

	return flags{configFile}
}
