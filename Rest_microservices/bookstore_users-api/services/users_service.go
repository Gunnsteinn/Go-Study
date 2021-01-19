package services

import (
	"github.com/Gunnsteinn/Go-Study/Rest_microservices/bookstore_users-api/domain/users"
	"github.com/Gunnsteinn/Go-Study/Rest_microservices/bookstore_users-api/utils/errors"
)

func CreateUSer(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
