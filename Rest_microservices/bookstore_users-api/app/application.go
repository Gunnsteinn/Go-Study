package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

// StartApplication is a func to map and route
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
