package main

func main() {
	flags := initFlags()
	config := ParseConfigFromFile(flags.configFile)
	notificationCallback := func() {
		notify(config.Messages)
	}
	StartTimer(config, notificationCallback)
}
