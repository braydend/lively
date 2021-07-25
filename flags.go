package main

import "flag"

type flags struct {
	configFile string
}

func initFlags() flags {
	var configFile string
	flag.StringVar(&configFile, "config", "", "A YAML file to configure the app")
	flag.Parse()

	return flags{configFile}
}
