package basket

type ResponseBasket struct {
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
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Amount uint   `json:"amount"`
}

type RequestProduct []ResponseProduct
