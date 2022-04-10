package basket

import (
	"errors"
	"fmt"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/order"
	"gorm.io/gorm"
)

type RepositoryBasket struct {
	db *gorm.DB
}

func NewRepositoryBasket(db *gorm.DB) *RepositoryBasket {
	return &RepositoryBasket{
		db: db,
	}
}

func (r *RepositoryBasket) Migration() {
	r.db.AutoMigrate(&order.Order{})
}

func (r *RepositoryBasket) GetAll() []order.Order {
	var baskets []order.Order

	r.db.Where("IsOrder =?", false).Find(&baskets)

	return baskets
}

func (r *RepositoryBasket) GetByID(id int) order.Order {
	var basket order.Order
	result := r.db.Where("IsOrder =?", false).Find(&basket, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Basket item not found with id : %d", id)
		return order.Order{}
	}
	return basket
}

func (r *RepositoryBasket) GetByCustomerID(id int) []order.Order {
	var baskets []order.Order
	result := r.db.Where("CustomerID = ?", id).Where("IsOrder =?", false).Find(&baskets)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Order not found with customer id : %d", id)
		return nil
	}
	return baskets
}

func (r *RepositoryBasket) GetByProductID(id int) []order.Order {
	var baskets []order.Order
	result := r.db.Where("ProductID = ?", id).Where("IsOrder =?", false).Find(&baskets)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Order not found with product id : %d", id)
		return nil
	}
	return baskets
}

func (r *RepositoryBasket) GetByIDsAndIsOrder(basket order.Order) order.Order {
	var basketItem order.Order
	result := r.db.Where("CustomerID = ?", basket.CustomerID).Where("ProductID = ?", basket.ProductID).Where("IsOrder =?", false).First(&basketItem)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return order.Order{}
	}
	return basketItem
}

func (r *RepositoryBasket) Delete(b order.Order) error {
	result := r.db.Delete(&b)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryBasket) DeleteByID(id int) error {
	result := r.db.Delete(&order.Order{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryBasket) Create(b *order.Order) error {
	result := r.db.Create(b)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryBasket) Update(b order.Order) error {
	result := r.db.Save(&b)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
