package main

import "flag"

type flags struct {
	interval int
}

func initFlags() flags {
	var interval int
	flag.IntVar(&interval, "interval", 30, "The amount of minutes between activity reminders")
	flag.Parse()

	return flags{interval}
}
