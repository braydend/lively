package main

func main() {
	flags := initFlags()
	config := ParseConfig(ReadFile(flags.configFile))
	notificationCallback := func() {
		notify(config.Messages)
	}
	StartTimer(config, notificationCallback)
}
