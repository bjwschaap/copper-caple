package caple

import "github.com/kataras/iris"

// statusHandler returns the current (health)status of the service
func statusHandler(c *iris.Context) {
	c.JSON(iris.StatusOK, iris.Map{
		"status": "OK",
	})
}
