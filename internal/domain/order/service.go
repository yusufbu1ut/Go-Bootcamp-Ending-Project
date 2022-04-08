package order

type ServiceOrder struct {
	r *RepositoryOrder
}

func NewServiceOrder(r *RepositoryOrder) *ServiceOrder {
	return &ServiceOrder{
		r: r,
	}
}

func (s *ServiceOrder) GetAll(pageIndex, pageSize int) ([]Order, int) {
	orders, count := s.r.GetAll(pageIndex, pageSize)
	return orders, count
}
