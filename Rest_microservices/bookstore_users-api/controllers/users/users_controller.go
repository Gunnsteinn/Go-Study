package users

import (
	"github.com/Gunnsteinn/Go-Study/Rest_microservices/bookstore_users-api/domain/users"
	"github.com/Gunnsteinn/Go-Study/Rest_microservices/bookstore_users-api/services"
	"github.com/Gunnsteinn/Go-Study/Rest_microservices/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body.")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUSer(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
