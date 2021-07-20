package lively

import (
	"fmt"
	"time"
)

func buildDuration(minutes int) time.Duration {
	duration, err := time.ParseDuration(fmt.Sprintf("%dm", minutes))

	if err != nil {
		panic(err)
	}

	return duration
}

func setTimer(minutes int) *time.Timer {
	return time.NewTimer(buildDuration(minutes))
}
