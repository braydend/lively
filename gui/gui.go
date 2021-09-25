package main

import (
	_ "embed"
	"flag"
	"github.com/wailsapp/wails"
)

//go:embed frontend/build/static/js/main.js
var js string

//go:embed frontend/build/static/css/main.css
var css string

var message string

func getMessage() string {
	//return func() string {
	return message
	//}
}

func main() {
	flag.StringVar(&message, "message", "", "Message to show in GUI")

	flag.Parse()

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "gui",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(getMessage)
	app.Run()
}