package caple

import (
	"strings"

	"github.com/kataras/iris"
)

type apiTokenMiddleware struct{}

func (m apiTokenMiddleware) Serve(c *iris.Context) {
	authHeader := c.RequestHeader("Authorization")
	if authHeader != "" {
		params := strings.Split(authHeader, " ")
		if params[0] == "apikey" && params[1] == config.apiKey {
			c.Next()
		} else {
			c.JSON(iris.StatusForbidden, iris.Map{
				"403": "Forbidden",
			})
		}
	} else {
		c.JSON(iris.StatusForbidden, iris.Map{
			"403": "Forbidden",
		})
	}
}
