package user

import (
	"errors"
	"fmt"
)

type QueryError struct {
	UserNotFound error
	UserNotExist error
	UserExist    error
}

func NewQueryError(u User) QueryError {
	return QueryError{
		UserNotFound: errors.New("no user founded"),
		UserNotExist: fmt.Errorf(fmt.Sprintf("user id: %s not exist", u.Id)),
		UserExist:    fmt.Errorf(fmt.Sprintf("user %s already exist", u.Name)),
	}
}
