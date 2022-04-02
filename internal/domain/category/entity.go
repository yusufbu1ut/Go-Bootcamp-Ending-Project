package category

import (
	"fmt"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string  `json:"name"`
	Height      float64 `json:"height"`
	Color       string  `json:"color"`
	Brand       string  `json:"brand"`
	Description string  `json:"description"`
}

func NewCategory(name string, height float64, color string, brand string, description string) *Category {
	return &Category{
		Name:        name,
		Height:      height,
		Color:       color,
		Brand:       brand,
		Description: description,
	}
}

func (c *Category) ToString() string {
	return fmt.Sprintf("Id: %d, Name: %s, Height: %f, Color: %s, Brand: %s", c.ID, c.Name, c.Height, c.Color, c.Brand)
}

func (c *Category) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Printf("Categoty %s deleting...", c.Name)
	return nil
}
