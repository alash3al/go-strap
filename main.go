package main

import (
	"core.globals"
	"core.models"
	"core.routes"
)

func main() {
	// the version number
	globals.LogSuccess("⇨ Current version is: version-%s \n", globals.Version)

	// load the global .env file
	globals.Boot()

	// connect to the mysql server
	models.Boot()

	// initalizing the routes server
	routes.Serve()
}
