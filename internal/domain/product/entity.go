package product

import (
	"fmt"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/category"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Name        string            `json:"name"`
	Price       float64           `json:"price"`
	Amount      uint              `json:"amount"`
	Code        uint              `json:"code"`
	Description string            `json:"description"`
	CategoryID  uint              `json:"category-id"`
	Category    category.Category `gorm:"foreignKey:CategoryID"`
}

func NewProduct(name string, price float64, amount uint, code uint, description string, category uint) *Product {
	return &Product{
		Name:        name,
		Price:       price,
		Amount:      amount,
		Code:        code,
		Description: description,
		CategoryID:  category,
	}
}

func (p *Product) ToString() string {
	return fmt.Sprintf("Code: %d, Name: %s, Amount: %d, Price: %f", p.Code, p.Name, p.Amount, p.Price)
}

func (p *Product) BeforeDelete(tx *gorm.DB) (err error) {
	log.Printf("Product %d %s %d deleting...", p.Code, p.Name, p.CategoryID)
	return nil
}
