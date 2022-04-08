package product

type ServiceProduct struct {
	r *RepositoryProduct
}

func NewServiceProduct(r *RepositoryProduct) *ServiceProduct {
	return &ServiceProduct{
		r: r,
	}
}

func (s *ServiceProduct) Create(product *Product) error {
	if product.Name == "" {
		return ErrProductNameNil
	}
	if product.Code <= 0 {
		return ErrProductCode
	}
	item := s.r.GetByCodeAndCategoryID(int(product.Code), int(product.CategoryID))
	if item.ID != 0 {
		return ErrProductWithExistCode
	}
	item = s.r.GetByFullNameAndCategoryID(product.Name, int(product.CategoryID))
	if item.ID != 0 {
		return ErrProductWithExistName
	}
	err := s.r.Create(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceProduct) GetAll(pageIndex, pageSize int) ([]Product, int) {
	products, count := s.r.GetAll(pageIndex, pageSize)
	return products, count
}

func (s *ServiceProduct) Update(product *Product) error {
	if product.ID == 0 && product.Name == "" && product.CategoryID == 0 && product.Code == 0 {
		return ErrProductFieldsRequested
	}
	item := s.r.GetByID(int(product.ID))
	if item.ID == 0 {
		return ErrProductNotExist
	}
	err := s.r.Update(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceProduct) Delete(product *Product) error {
	productItem := s.r.Get(product)
	if productItem.ID == 0 {
		return ErrProductNotExist
	}
	err := s.r.Delete(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceProduct) DeleteWithID(id int) error {
	product := s.r.GetByID(id)
	if product.ID == 0 {
		return ErrProductNotExist
	}
	err := s.r.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceProduct) Search(product *Product) ([]Product, error) {
	var products []Product
	if product.CategoryID != 0 && product.Amount == 0 && product.Name == "" {
		products = s.r.GetByCategoryID(int(product.CategoryID))
	} else if product.Amount != 0 && product.CategoryID == 0 && product.Name == "" {
		products = s.r.GetByAmount(int(product.Amount))
	} else if product.Name != "" && product.CategoryID != 0 && product.Amount == 0 {
		products = s.r.GetByNameAndCategoryID(product.Name, int(product.CategoryID))
	} else if product.Name != "" && product.CategoryID == 0 && product.Amount != 0 {
		products = s.r.GetByNameAndAmount(product.Name, int(product.Amount))
	} else if product.Name == "" && product.CategoryID != 0 && product.Amount != 0 {
		products = s.r.GetByCategoryIDAndAmount(int(product.CategoryID), int(product.Amount))
	} else if product.Name != "" && product.CategoryID != 0 && product.Amount != 0 {
		products = s.r.GetByNameAndCategoryIDAndAmount(product.Name, int(product.CategoryID), int(product.Amount))
	} else {
		products = s.r.GetByName(product.Name)
	}

	return products, nil
}
