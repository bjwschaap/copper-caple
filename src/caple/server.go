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
	api.Use(logger.New())
	api.Use(&apiTokenMiddleware{})
	api.Get("/caple/v1/status", statusHandler)
	api.ListenTLS(config.listenAddress, "server.cert", "server.key")

	return nil
}

// statusHandler returns the current (health)status of the service
func statusHandler(c *iris.Context) {
	c.JSON(iris.StatusOK, iris.Map{
		"status": "OK",
	})
}
