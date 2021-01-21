package mysql_utils

import (
	"github.com/Gunnsteinn/Go-Study/Rest_microservices/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	indexUniqueEmail = 1062
	errorNoRows      = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	// Assertion with two return values.
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id.")
		}
		return errors.NewInternalServerError("error parsing database response.")
	}

	switch sqlErr.Number {
	case indexUniqueEmail:
		return errors.NewBadRequestError("invalid data.")
	}
	return errors.NewInternalServerError("error processing request.")
}