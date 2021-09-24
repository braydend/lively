package main

import (
	"log"
	"os/exec"
)

func main() {
	flags := initFlags()
	config := ParseConfigFromFile(flags.configFile)
	notificationCallback := func() {
		notify := func() {
			cmd := exec.Command("gui/build/gui", "--message", GetRandomMessage(config.Messages))
			err := cmd.Run()

			if err != nil {
				log.Fatalln(err)
			}
		}

		go notify()
	}
	StartTimer(config, notificationCallback)
}
