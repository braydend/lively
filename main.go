package main

func main() {
	flags := initFlags()
	config := ParseConfig(ReadFile(flags.configFile))
	StartTimer(config, notify)
}
