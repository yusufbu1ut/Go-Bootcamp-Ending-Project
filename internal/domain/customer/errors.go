package customer

import (
	"errors"
)

var (
	ErrCustomerExistWithUsername = errors.New("Customer already exist with same Username in database")
	ErrCustomerExistWithEmail    = errors.New("Customer already exist with same Email in database")
	ErrCustomerUsernameNil       = errors.New("Customer username can not be default")
	ErrCustomerEmailNil          = errors.New("Customer email can not be default")
	ErrCustomerPasswordNil       = errors.New("Customer password can not be default")
)
