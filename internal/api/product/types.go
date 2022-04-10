package product

type ResponseProduct struct {
	ID          uint    `json:"id"`
	Code        uint    `json:"code"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Amount      uint    `json:"amount"`
	CategoryId  uint    `json:"category-id"`
	Description string  `json:"description"`
}
