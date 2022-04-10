package sign_up

type RequestCustomer struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	PhoneNo  string `json:"phone"`
	Address  string `json:"address"`
}
