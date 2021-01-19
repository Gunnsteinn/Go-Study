package app

import (
	"github.com/Gunnsteinn/Go-Study/Rest_microservices/bookstore_users-api/controllers/ping"
	"github.com/Gunnsteinn/Go-Study/Rest_microservices/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", users.SearchUSer)
	router.POST("/users", users.CreateUser)

}
