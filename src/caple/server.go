package caple

import (
	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
	"github.com/urfave/cli"
)

// StartServer starts the REST Service using Iris
func StartServer(c *cli.Context) error {
	// Store configuration globally
	setConfig(c)

	// Start server, and serve context path
	api := iris.New()

	// Set some configuration
	api.Config.LoggerPreffix = "[copper-caple] "

	// Add middlewares that process all requests before the handler
	api.Use(logger.New())
	api.Use(&apiTokenMiddleware{})

	// Add all API endpoint handlers
	api.Get("/caple/v1/status", statusHandler)

	// Start listening..
	api.ListenTLS(config.listenAddress, "server.cert", "server.key")

	// Eventually return to CLI wrapper
	return nil
}

// statusHandler returns the current (health)status of the service
func statusHandler(c *iris.Context) {
	c.JSON(iris.StatusOK, iris.Map{
		"status": "OK",
	})
}
