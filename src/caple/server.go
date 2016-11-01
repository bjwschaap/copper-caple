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

	// Initialize DB connection
	dbConnect()

	// Start server, and serve context path
	api := iris.New()

	// Set some configuration
	api.Config.LoggerPreffix = "[copper-caple] "

	// Add middlewares that process all requests before the handler
	api.Use(logger.New())
	api.Use(&apiTokenMiddleware{})

	// Add all API endpoint handlers
	api.Get("/caple/v1/status", statusHandler)
	api.Get("/caple/v1/students", studentsHandler)
	api.Get("/caple/v1/student/:id", studentByIDHandler)

	// Start listening..
	api.ListenTLS(config.listenAddress, "server.cert", "server.key")

	// Eventually return to CLI wrapper
	return nil
}
