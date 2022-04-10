package order

import (
	"fmt"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/customer"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/product"
	"gorm.io/gorm"
	"log"
)

type Order struct {
	gorm.Model
	CustomerID uint `json:"customer-id"`
	ProductID  uint `json:"product-id"`
	Amount     uint `json:"amount"`
	IsOrder    bool
	OrderCode  string
	Customer   customer.Customer `gorm:"foreignKey:CustomerID"`
	Product    product.Product   `gorm:"foreignKey:ProductID"`
}

func NewOrder(customer uint, product uint, amount uint) *Order {
	return &Order{
		CustomerID: customer,
		ProductID:  product,
		Amount:     amount,
		IsOrder:    false,
	}
}

func (o *Order) ToString() string {
	return fmt.Sprintf("CustomerID: %d, ProductID: %d, Amount: %d ", o.CustomerID, o.ProductID, o.Amount)
}

func (o *Order) BeforeDelete(tx *gorm.DB) (err error) {
	log.Printf("Order %d deleting...", o.ID)
	return nil
}
