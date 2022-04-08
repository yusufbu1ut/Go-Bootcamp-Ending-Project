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

	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products)
	r.db.Model(&Product{}).Count(&count)

	return products, int(count)
}

func (r *RepositoryProduct) Get(product *Product) Product {
	var productItem Product
	result := r.db.Where(&Product{
		Name:       product.Name,
		Code:       product.Code,
		CategoryID: product.CategoryID,
	}).First(&productItem)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Product not found ")
		return Product{}
	}
	return productItem
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
		fmt.Printf("Product not found with code : %d", code)
		return Product{}
	}
	return product
}

func (r *RepositoryProduct) GetByFullName(name string) Product {
	var product Product
	result := r.db.Where("Name=?", name).First(product)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Product not found with code : %s", name)
		return Product{}
	}
	return product
}

func (r *RepositoryProduct) GetByAmount(amount int) []Product {
	var products []Product
	r.db.Where("Amount >= ?", amount).Find(&products)
	return products
}

func (r *RepositoryProduct) GetByName(name string) []Product {
	var products []Product
	r.db.Where("Name LIKE ?", "%"+name+"%").Find(&products)
	return products
}

func (r *RepositoryProduct) GetByCategoryID(id int) []Product {
	var products []Product
	r.db.Where("CategoryID=?", id).Find(&products)
	return products
}

func (r *RepositoryProduct) GetByNameAndCategoryID(name string, id int) []Product {
	var products []Product
	r.db.Where("Name LIKE ?", "%"+name+"%").Where("CategoryID = ?", id).Find(&products)
	return products
}

func (r *RepositoryProduct) GetByNameAndAmount(name string, amount int) []Product {
	var products []Product
	r.db.Where("Name LIKE ?", "%"+name+"%").Where("Amount >= ?", amount).Find(&products)
	return products
}

func (r *RepositoryProduct) GetByFullNameAndCategoryID(name string, id int) Product {
	var product Product
	result := r.db.Where(&Product{Name: name, CategoryID: uint(id)}).First(&product)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Product not found with code & id : %s , %d", name, id)
		return Product{}
	}
	return product
}

func (r *RepositoryProduct) GetByCodeAndCategoryID(code int, id int) Product {
	var product Product
	result := r.db.Where(&Product{Code: uint(code), CategoryID: uint(id)}).First(&product)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Product not found with code & id : %d , %d", code, id)
		return Product{}
	}
	return product
}

func (r *RepositoryProduct) GetByCategoryIDAndAmount(id int, amount int) []Product {
	var products []Product
	r.db.Where(&Product{CategoryID: uint(id), Amount: uint(amount)}).Find(&products)
	return products
}

func (r *RepositoryProduct) GetByNameAndCategoryIDAndAmount(name string, id int, amount int) []Product {
	var products []Product
	r.db.Where(&Product{Name: name, CategoryID: uint(id), Amount: uint(amount)}).Find(&products)
	return products
}

func (r *RepositoryProduct) Delete(p *Product) error {
	product := r.Get(p)
	result := r.db.Delete(&Product{}, product.ID)
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

func (r *RepositoryProduct) Update(p *Product) error {
	result := r.db.Save(p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
