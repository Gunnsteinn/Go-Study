package app

import (
	"github.com/Gunnsteinn/Go-Study/Rest_microservices/bookstore_users-api/controllers/ping"
	"github.com/Gunnsteinn/Go-Study/Rest_microservices/bookstore_users-api/controllers/users"
)

// mapUrls is used to mapping urls
func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
}
