package caple

import (
	"log"

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
	api.Config.LoggerPreffix = loggerPrefix

	// Add middlewares that process all requests alongside the handler
	api.Use(logger.New())
	api.Use(&apiTokenMiddleware{})

	// Add all API endpoint handlers
	api.Get(contextPath+"/status", statusHandler)
	api.Get(contextPath+"/students", studentsHandler)
	api.Get(contextPath+"/student/:id", studentByIDHandler)

	// Optionally output service endpoint
	if config.debug {
		log.Printf("Service URL: %s\n", "https://"+config.listenAddress+contextPath)
	}

	// Start listening..
	api.ListenTLS(config.listenAddress, "server.cert", "server.key")

	// Eventually return to CLI wrapper
	return nil
}
