package services

import (
	"github.com/Gunnsteinn/Go-Study/Rest_microservices/bookstore_users-api/domain/users"
	"github.com/Gunnsteinn/Go-Study/Rest_microservices/bookstore_users-api/utils/crypto_utils"
	"github.com/Gunnsteinn/Go-Study/Rest_microservices/bookstore_users-api/utils/date_utils"
	"github.com/Gunnsteinn/Go-Study/Rest_microservices/bookstore_users-api/utils/errors"
)

const (
	emptyValue = ""
)

// GetUser is the business logic to get user from the db.
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateUser is the business logic to create a new user from the database.
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser is the business logic to update an exist user from the database.
func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != emptyValue {
			current.FirstName = user.FirstName
		}

		if user.LastName != emptyValue {
			current.LastName = user.LastName
		}

		if user.Email != emptyValue {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

// DeleteUser is the business logic to delete an exist user from the database.
func DeleteUser(userID int64) *errors.RestErr {
	user := &users.User{ID: userID}
	return user.Delete()
}

//
func Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{Status: status}
	return dao.FindByStatus(status)
}
