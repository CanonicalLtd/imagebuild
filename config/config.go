package config

import "os"

// Default settings
const (
	DefaultPort          = ":8000"
	DefaultDocRoot       = "webapp/build"
	DefaultIndexTemplate = "index.html"
)

// Settings defines the application configuration
type Settings struct {
	Port          string
	DocRoot       string
	IndexTemplate string
}

// ParseArgs checks the environment variables
func ParseArgs() *Settings {
	var (
		port = DefaultPort
	)

	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	return &Settings{
		Port:          port,
		DocRoot:       DefaultDocRoot,
		IndexTemplate: DefaultIndexTemplate,
	}
}
