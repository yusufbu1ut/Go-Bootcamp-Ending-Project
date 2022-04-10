package order

import (
	"time"
)

type ResponseOrder struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	PhoneNo  string `json:"phone"`
	Address  string `json:"address"`
	Products []ResponseProduct
}

type ResponseProduct struct {
	OrderTime time.Time `json:"order-time"`
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Amount    uint      `json:"amount"`
	Code      string    `json:"code"`
}
