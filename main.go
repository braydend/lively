package main

func main() {
	flags := initFlags()
	timer := setTimer(flags.interval)

	for {
		select {
			case <-timer.C:
				notify()
				timer = setTimer(flags.interval)
			}
	}
}
