package main

import (
	"github.com/writetosalman/go-rest-api-boilerplate/core/server"
	"github.com/writetosalman/go-rest-api-boilerplate/config"
)

// main function
func main() {

	// Initialise .env file
	config.Initialise()

	// Start Server
	server.StartServer()
}

