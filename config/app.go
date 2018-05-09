package config


import (
	"os"
	"github.com/writetosalman/go-rest-api-boilerplate/utilities"
	"github.com/subosito/gotenv"
)

// This function initialises ENV file
func Initialise() {
	gotenv.Load()
}

// This function returns environment variables
func Getenv(key string) string {

	envVar := os.Getenv(key)
	if envVar != "" {
		return envVar
	}

	// Fatal error
	utilities.Log("Critical error. Environment variable ["+ key +"] not found")
	os.Exit(500)
	return ""
}