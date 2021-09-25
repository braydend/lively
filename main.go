package main

import (
	"lively/utils"
	"log"
	"os/exec"
)

func main() {
	flags := utils.InitFlags()
	config := utils.ParseConfigFromFile(flags.ConfigFile)
	notificationCallback := func() {
		notify := func() {
			cmd := exec.Command("gui/build/gui", "--message", utils.GetRandomMessage(config.Messages))
			err := cmd.Run()

			if err != nil {
				log.Fatalln(err)
			}
		}

		go notify()
	}
	utils.StartTimer(config, notificationCallback)
}
