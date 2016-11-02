package caple

import (
	"time"

	"github.com/kataras/iris"
	pg "gopkg.in/pg.v5"
)

// statusHandler returns the current (health)status of the service
func statusHandler(c *iris.Context) {
	var num int
	start := time.Now()
	_, err := db.Query(pg.Scan(&num), "SELECT ?", 42)
	elapsed := time.Since(start)
	if config.debug {
		c.Log("Status Query took: %s", elapsed)
	}
	if err == nil && num == 42 {
		c.JSON(iris.StatusOK, iris.Map{
			"status": "OK",
		})
	} else {
		c.JSON(iris.StatusServiceUnavailable, iris.Map{
			"error": err.Error(),
		})
	}
}
