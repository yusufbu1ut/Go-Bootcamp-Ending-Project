package product

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type RepositoryProduct struct {
	db *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *RepositoryProduct {
	return &RepositoryProduct{
		db: db,
	}
}

func (r *RepositoryProduct) Migration() {
	r.db.AutoMigrate(&Product{})
}

func (r *RepositoryProduct) InsertSampleData(products []Product) {
	for _, c := range products {
		r.Create(&c)
	}
}

func (r *RepositoryProduct) GetAll(pageIndex, pageSize int) ([]Product, int) {
	var products []Product
	var count int64

	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)

	return products, int(count)
}

func (r *RepositoryProduct) GetByID(id int) Product {
	var product Product
	result := r.db.Find(&product, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Category not found with id : %d", id)
		return Product{}
	}
	return product
}

func (r *RepositoryProduct) GetByCode(code int) Product {
	var product Product
	result := r.db.Where("Code=?", code).First(product)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Category not found with code : %d", code)
		return Product{}
	}
	return product
}

func (r *RepositoryProduct) GetByName(name string) []Product {
	var products []Product
	r.db.Where("Name LIKE ?", "%"+name+"%").Find(&products)
	return products
}

func (r *RepositoryProduct) GetByCategoryID(id int) []Product {
	var products []Product
	r.db.Where("CategoryID=?", id).Find(products)
	return products
}

func (r *RepositoryProduct) Delete(p Product) error {
	result := r.db.Delete(p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryProduct) DeleteByID(id int) error {
	result := r.db.Delete(&Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryProduct) Create(p *Product) error {
	result := r.db.Create(p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryProduct) Update(p Product) error {
	result := r.db.Save(p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
