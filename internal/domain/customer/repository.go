package customer

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type RepositoryCustomer struct {
	db *gorm.DB
}

func NewRepositoryCustomer(db *gorm.DB) *RepositoryCustomer {
	return &RepositoryCustomer{
		db: db,
	}
}

func (r *RepositoryCustomer) Migration() {
	r.db.AutoMigrate(&Customer{})
}

func (r *RepositoryCustomer) InsertSampleData(customers []Customer) {
	for _, c := range customers {
		r.Create(&c)
	}
}

func (r *RepositoryCustomer) GetAll(pageIndex, pageSize int) ([]Customer, int) {
	var customers []Customer
	var count int64

	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&customers).Count(&count)

	return customers, int(count)
}

func (r *RepositoryCustomer) GetByID(id int) Customer {
	var customer Customer
	result := r.db.Find(&customer, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Customer not found with id : %d", id)
		return Customer{}
	}
	return customer
}

func (r *RepositoryCustomer) GetByName(name string) []Customer {
	var customers []Customer
	r.db.Where("Name LIKE ?", "%"+name+"%").Find(&customers)
	return customers
}

func (r *RepositoryCustomer) GetByCustomerName(name string) []Customer {
	var customers []Customer
	r.db.Where("CustomerName LIKE ?", "%"+name+"%").Find(&customers)
	return customers
}

func (r *RepositoryCustomer) GetByMail(mail string) []Customer {
	var customers []Customer
	r.db.Where("Email = ?", mail).Find(&customers)
	return customers
}

func (r *RepositoryCustomer) GetByMailAndPassword(mail string, pass string) []Customer {
	var customers []Customer
	r.db.Raw("SELECT * FROM Customer WHERE Email = ? AND Password =?", mail, pass).Scan(&customers)
	return customers
}

func (r *RepositoryCustomer) Delete(c Customer) error {
	result := r.db.Delete(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryCustomer) DeleteByID(id int) error {
	result := r.db.Delete(&Customer{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryCustomer) Create(c *Customer) error {
	result := r.db.Create(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryCustomer) Update(c Customer) error {
	result := r.db.Save(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
