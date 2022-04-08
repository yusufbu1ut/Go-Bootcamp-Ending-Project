package order

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type RepositoryOrder struct {
	db *gorm.DB
}

func NewRepositoryOrder(db *gorm.DB) *RepositoryOrder {
	return &RepositoryOrder{
		db: db,
	}
}

func (r *RepositoryOrder) Migration() {
	r.db.AutoMigrate(&Order{})
}

func (r *RepositoryOrder) InsertSampleData(orders []Order) {
	for _, o := range orders {
		r.Create(&o)
	}
}

func (r *RepositoryOrder) GetAll(pageIndex, pageSize int) ([]Order, int) {
	var orders []Order
	var count int64

	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&orders)
	r.db.Model(&Order{}).Count(&count)
	return orders, int(count)
}

func (r *RepositoryOrder) GetByID(id int) Order {
	var order Order
	result := r.db.Find(&order, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Order not found with id : %d", id)
		return Order{}
	}
	return order
}

func (r *RepositoryOrder) GetByCustomerID(id int) []Order {
	var orders []Order
	result := r.db.Where("CustomerID = ?", id).Find(&orders)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Order not found with customer id : %d", id)
		return nil
	}
	return orders
}

func (r *RepositoryOrder) GetByProductID(id int) []Order {
	var orders []Order
	result := r.db.Where("ProductID = ?", id).Find(&orders)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Order not found with product id : %d", id)
		return nil
	}
	return orders
}

func (r *RepositoryOrder) Delete(o Order) error {
	result := r.db.Delete(o)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryOrder) DeleteByID(id int) error {
	result := r.db.Delete(&Order{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryOrder) Create(o *Order) error {
	result := r.db.Create(o)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryOrder) Update(o Order) error {
	result := r.db.Save(o)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
