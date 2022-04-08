package category

import (
	"errors"
)

var (
	ErrCategoryExistWithCode = errors.New("Category already exist with same code in database")
	ErrCategoryExistWithName = errors.New("Category already exist with same name in database")
	ErrCategoryNameNil       = errors.New("Category name can not be default")
	ErrCategoryCodeZero      = errors.New("Category code can not be default")
)
