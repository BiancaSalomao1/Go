package domain

import (
	"errors"
	"fmt"
)


var (
	ErrProductNotFound = errors.New("product not found")
	ErrOrderNotFound = errors.New("order not found")
	ErrProductAlreadyInOrder = errors.New("product already in order")
	ErrProductQuantityInvalid = errors.New("product quantity invalid")
	ErrProductNotEnoughStock = errors.New("product not enough stock")
	ErrCustomerNotFound = errors.New("customer not found")
	ErrCustomerInvalid = errors.New("customer invalid")
	ErrChangeStatusInvalid = errors.New("change status invalid")
)



func IsError(err error, target error) bool {
	return errors.Is(err, target)
}

func Errorf(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}