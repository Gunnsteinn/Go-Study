package users

import (
	"github.com/Gunnsteinn/Go-Study/Rest_microservices/bookstore_users-api/domain/users"
	"github.com/Gunnsteinn/Go-Study/Rest_microservices/bookstore_users-api/services"
	"github.com/Gunnsteinn/Go-Study/Rest_microservices/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//
func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number.")
	}
	return userID, nil
}

// Create create a new user.
func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body.")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// Get get an existing user.
func Get(c *gin.Context) {
	userID, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.UsersService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

// Update update an exist user.
func Update(c *gin.Context) {
	userID, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body.")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.ID = userID

	// Decide if it's a full update (PUT) or a partial update (PATCH).
	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UsersService.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// Delete delete an exist user.
func Delete(c *gin.Context) {
	userID, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if deleteErr := services.UsersService.DeleteUser(userID); deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Search find all the user by status
func Search(c *gin.Context) {
	status := c.Query("status")

	users, searchErr := services.UsersService.Search(status)
	if searchErr != nil {
		c.JSON(searchErr.Status, searchErr)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}
