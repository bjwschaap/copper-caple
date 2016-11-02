package caple

import (
	"log"
	"reflect"

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

	// Set some logging configuration
	api.Config.LoggerPreffix = loggerPrefix
	log.SetPrefix(loggerPrefix)

	// Add middlewares that process all requests alongside the handler
	api.Use(logger.New())
	api.Use(&apiTokenMiddleware{})

	// Add all API endpoint handlers
	api.Get(contextPath+"/status", statusHandler)
	api.Get(contextPath+"/students", studentsHandler)
	api.Get(contextPath+"/student/:id", studentByIDHandler)

	// Optionally output service endpoint
	if config.debug {
		log.Printf("Base URL:    %s\n", "https://"+config.listenAddress+contextPath)
		log.Println("Registered endpoints:")
		inspectRoutes(api)
	}

	// Start listening..
	api.ListenTLS(config.listenAddress, "server.cert", "server.key")

	// Eventually return to CLI wrapper
	return nil
}

// inspectRoutes uses reflection to extract route information from the
// Iris framework. Reflection is needed because the routes are unexported from
// the iris package.
func inspectRoutes(api *iris.Framework) {
	// Get muxAPI field by reflection
	muxV := reflect.ValueOf(api).Elem().FieldByName("muxAPI").Elem()
	// Get apiRoutes field by reflection
	routesV := muxV.FieldByName("apiRoutes")
	// Extract slice with routes
	routes := routesV.Slice(0, routesV.Len()-1)
	for i := 0; i < routes.Len(); i++ {
		// Get each route by reflection
		rv := routes.Index(i).Elem()
		// Extract METHOD and PATH by reflection/inspection
		methodStr := rv.FieldByName("methodStr").String()
		path := rv.FieldByName("path").String()
		log.Printf("%d - %s %s", i+1, methodStr, path)
	}
}
