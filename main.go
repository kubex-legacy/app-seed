package main

import (
	"github.com/kubex/potens-go"
)

var (
	app *potens.Application
)

func main() {
	app = potens.QuickStartApp("Scratch App")
	startup()
	app.RegisterAsAppServer()
	app.FatalErr(app.ExposeAndServe())
}

func startup() {
	s := app.AppServer()
	//Setup Routes Here
	s.HandlePage("/", PageDefinition, DefaultPage)
}
