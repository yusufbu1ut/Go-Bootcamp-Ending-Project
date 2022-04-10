package category

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Category struct {
	gorm.Model
	Name        string `json:"name"`
	Code        uint   `json:"code"`
	Description string `json:"description"`
}

func NewCategory(name string, code uint, description string) *Category {
	return &Category{
		Name:        name,
		Code:        code,
		Description: description,
	}
}

func (c *Category) ToString() string {
	return fmt.Sprintf("Name: %s, Code: %d", c.Name, c.Code)
}

func (c *Category) BeforeDelete(tx *gorm.DB) (err error) {
	log.Printf("Categoty %s deleting...", c.Name)
	return nil
}
