package customer

import (
	"fmt"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name         string `json:"name"`
	CustomerName string `json:"customer-name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	PhoneNo      string `json:"phone"`
	Address      string `json:"address"`
	LogStatus    bool   `json:"status"`
}

func NewCustomer(name string, customerName string, email string, pass string, phone string, address string) *Customer {
	return &Customer{
		Name:         name,
		CustomerName: customerName,
		Email:        email,
		Password:     pass,
		PhoneNo:      phone,
		Address:      address,
	}
}

func (c *Customer) ToString() string {
	return fmt.Sprintf("Id: %d, Name: %s, Mail: %s, Phone: %s, Address: %s", c.ID, c.Name, c.Email, c.PhoneNo, c.Address)
}

func (c *Customer) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Printf("Customer %s deleting...", c.Email)
	return nil
}
