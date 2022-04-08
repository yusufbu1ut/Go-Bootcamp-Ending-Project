package product

import "errors"

var (
	ErrProductWithExistName   = errors.New("Product already exist with same name in database")
	ErrProductWithExistCode   = errors.New("Product already exist with same code in database")
	ErrProductNameNil         = errors.New("Product name can not be empty")
	ErrProductCode            = errors.New("Product code can not be zero or negative")
	ErrProductFieldsRequested = errors.New("Product fields are requested to update")
	ErrProductNotExist        = errors.New("Product is not exist")
)
