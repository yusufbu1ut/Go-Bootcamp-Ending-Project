package category

import (
	"log"
)

type ServiceCategory struct {
	r *RepositoryCategory
}

func NewServiceCategory(r *RepositoryCategory) *ServiceCategory {
	return &ServiceCategory{
		r: r,
	}
}

func (s *ServiceCategory) Create(category *Category) error {
	if category.Name == "" {
		return ErrCategoryNameNil
	}
	if category.Code == 0 {
		return ErrCategoryCodeZero
	}
	existCats := s.r.GetByName(category.Name)
	if existCats != nil && len(existCats) > 0 {
		return ErrCategoryExistWithName
	}
	existCat := s.r.GetByCode(int(category.Code))
	if existCat.ID != 0 {
		return ErrCategoryExistWithCode
	}
	err := s.r.Create(category)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceCategory) GetAll(pageIndex, pageSize int) ([]Category, int) {
	categories, count := s.r.GetAll(pageIndex, pageSize)
	return categories, count
}

func (s *ServiceCategory) CreateWithCollectedData(categories chan Category) error {
	for category := range categories {
		err := s.Create(&category)
		if err != nil {
			log.Printf("Err: %s \n Category not created: %s \n", err.Error(), category.ToString())
			// in here internal server err can be returned cause of db create proccess
		}
	}
	return nil
}

func (s *ServiceCategory) GetCategoryWithCode(code int) Category {
	category := s.r.GetByCode(code)
	if category.ID == 0 {
		log.Printf("message: Category not found with code %d", code)
	}
	return category
}

func (s *ServiceCategory) GetCategoryWithId(id int) Category {
	category := s.r.GetByID(id)
	if category.ID == 0 {
		log.Printf("message: Category not found with id %d", id)
		return Category{}
	}
	return category
}
