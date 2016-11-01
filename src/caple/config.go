package caple

import (
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
	proxy         string
	apiKey        string
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
	config.listenAddress = c.String("listen")
	config.apiKey = c.String("apikey")
	config.proxy = os.Getenv("http_proxy")
	if config.proxy == "" {
		// Try uppercase as well..
		config.proxy = os.Getenv("HTTP_PROXY")
	}
}
