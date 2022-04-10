package customer

import (
	"fmt"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/hashing"
	"gorm.io/gorm"
	"log"
)

type Customer struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	PhoneNo  string `json:"phone"`
	Address  string `json:"address"`
}

func NewCustomer(name string, Username string, email string, pass string, phone string, address string) *Customer {
	password, err := hashing.HashWord(pass)
	if err != nil {
		fmt.Println("Error occurred: ", err.Error())
	}
	return &Customer{
		Name:     name,
		Username: Username,
		Email:    email,
		Password: password,
		PhoneNo:  phone,
		Address:  address,
	}
}

func (c *Customer) ToString() string {
	return fmt.Sprintf("Id: %d, Name: %s, Mail: %s, Phone: %s, Address: %s", c.ID, c.Name, c.Email, c.PhoneNo, c.Address)
}

func (c *Customer) BeforeDelete(tx *gorm.DB) (err error) {
	log.Printf("Customer %s deleting...", c.Email)
	return nil
}
