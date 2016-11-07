package caple

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

// Configuration is a structure to hold global configuration settings
type configuration struct {
	listenAddress string
	dbURL         string
	dbUser        string
	dbPassword    string
	dbName        string
	dbPoolSize    int
	proxy         string
	apiKey        string
	debug         bool
}

// Global configuration for this service
var config configuration

// setConfig copies cli and environment parameters/settings/variables to
// the global configuration
func setConfig(c *cli.Context) {
	config.dbURL = c.String("address")
	config.dbUser = c.String("user")
	config.dbPassword = c.String("password")
	config.dbName = c.String("db")
	config.dbPoolSize = c.Int("poolsize")
	config.listenAddress = c.String("listen")
	config.apiKey = c.String("apikey")
	config.debug = c.Bool("debug")
	config.proxy = os.Getenv("http_proxy")
	if config.proxy == "" {
		// Try uppercase as well..
		config.proxy = os.Getenv("HTTP_PROXY")
	}

	if config.debug {
		log.SetPrefix(loggerPrefix)
		log.Printf("DB Address:  %s\n", config.dbURL)
		log.Printf("DB Username: %s\n", config.dbUser)
		log.Printf("DB Password: %s\n", "********")
		log.Printf("DB Name:     %s\n", config.dbName)
		log.Printf("DB Poolsize: %d\n", config.dbPoolSize)
		log.Printf("API Key:     %s\n", config.apiKey)
		log.Printf("Proxy:       %s\n", config.proxy)
	}
}
