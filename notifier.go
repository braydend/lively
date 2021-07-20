package main

import (
	"github.com/gen2brain/beeep"
	"math/rand"
	"time"
)

func notify() {
	defaultMessages := []string{"Time to move around!", "Go for a quick walk!", "One quick minute of exercise is all it takes!"}

	err := beeep.Notify("Lively", getMessage(defaultMessages), "")
	if err != nil {
		panic(err)
	}
}

func getMessage(messages []string) string {
	rand.Seed(time.Now().UnixNano())

	return messages[rand.Int() % len(messages)]
}
