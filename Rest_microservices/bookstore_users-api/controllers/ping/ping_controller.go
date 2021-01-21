package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping func monitors the status of your application.
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
