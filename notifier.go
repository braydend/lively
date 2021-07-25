package main

import (
	"github.com/gen2brain/beeep"
	"math/rand"
	"time"
)

func notify(messages []string) {
	err := beeep.Alert("Lively", getRandomMessage(messages), "")
	if err != nil {
		panic(err)
	}
}

func getRandomMessage(messages []string) string {
	rand.Seed(time.Now().UnixNano())

	return messages[rand.Int()%len(messages)]
}
