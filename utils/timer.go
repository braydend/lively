package utils

import (
	"fmt"
	"regexp"
	"strconv"
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

func convertToMinutes(hours, minutes int) int {
	return hours*60 + minutes
}

func convertTimeToMinutes(t time.Time) int {
	return convertToMinutes(t.Hour(), t.Minute())
}

func isTimeWithin(timeToCheck time.Time, start, end string) bool {
	hasStartTime := start != ""
	hasEndTime := end != ""

	if hasStartTime {
		if convertTimeToMinutes(timeToCheck) < convertToMinutes(getTime24H(start)) {
			return false
		}
	}

	if hasEndTime {
		if convertTimeToMinutes(timeToCheck) > convertToMinutes(getTime24H(end)) {
			return false
		}
	}

	return true
}

func getTime24H(time24H string) (hours, minutes int) {
	hhmmRegex := regexp.MustCompile("(0[0-9]|1[0-9]|2[0-3])([0-5][0-9])")

	if !hhmmRegex.Match([]byte(time24H)) {
		panic(fmt.Errorf("failed to parse 24H time from: %s", time24H))
	}

	hours, err := strconv.Atoi(time24H[0:2])

	if err != nil {
		panic(err)
	}

	minutes, err = strconv.Atoi(time24H[2:])

	if err != nil {
		panic(err)
	}

	return hours, minutes
}

func StartTimer(config Config, callback func()) {
	timer := setTimer(config.Interval)

	for {
		select {
		case <-timer.C:
			if isTimeWithin(time.Now(), config.Start, config.End) {
				callback()
			}
			timer = setTimer(config.Interval)
		}
	}
}
