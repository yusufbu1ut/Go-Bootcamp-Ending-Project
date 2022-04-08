package category

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type RepositoryCategory struct {
	db *gorm.DB
}

func NewRepositoryCategory(db *gorm.DB) *RepositoryCategory {
	return &RepositoryCategory{
		db: db,
	}
}

func (r *RepositoryCategory) Migration() {
	r.db.AutoMigrate(&Category{})
}

func (r *RepositoryCategory) InsertSampleData() {
	DefaultCategory := NewCategory("Other", 0, "If there is no category-id this connection releases.")
	categ := r.GetByCode(int(DefaultCategory.Code))
	if categ.ID == 0 {
		r.db.Create(DefaultCategory)
	}
}

func (r *RepositoryCategory) GetAll(pageIndex, pageSize int) ([]Category, int) {
	var categories []Category
	var count int64

	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&categories)
	r.db.Model(&Category{}).Count(&count)
	return categories, int(count)
}

func (r *RepositoryCategory) GetByID(id int) Category {
	var category Category
	result := r.db.First(&category, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Category not found with id : %d", id)
		return Category{}
	}
	return category
}

func (r *RepositoryCategory) GetByName(name string) []Category {
	var categories []Category
	r.db.Where("Name LIKE ?", "%"+name+"%").Find(&categories)
	return categories
}

func (r *RepositoryCategory) GetByCode(code int) Category {
	var category Category
	result := r.db.Where("Code=?", code).Find(&category)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Category not found with code : %d", code)
		return Category{}
	}
	return category
}

func (r *RepositoryCategory) Delete(c Category) error {
	result := r.db.Delete(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryCategory) DeleteByID(id int) error {
	result := r.db.Delete(&Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryCategory) Create(c *Category) error {
	result := r.db.Create(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryCategory) Update(c Category) error {
	result := r.db.Save(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
