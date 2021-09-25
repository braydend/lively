package utils

import (
	"math/rand"
	"time"
)

func GetRandomMessage(messages []string) string {
	rand.Seed(time.Now().UnixNano())

	return messages[rand.Int()%len(messages)]
}
