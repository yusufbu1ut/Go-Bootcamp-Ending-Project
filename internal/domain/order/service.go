package order

import "time"

type ServiceOrder struct {
	r *RepositoryOrder
}

func NewServiceOrder(r *RepositoryOrder) *ServiceOrder {
	return &ServiceOrder{
		r: r,
	}
}

func (s *ServiceOrder) GetAll(id int) []Order {
	orders := s.r.GetByCustomerID(id)
	return orders
}

func (s *ServiceOrder) GetWithCode(id int, code string) []Order {
	orders := s.r.GetByOrderCode(id, code)
	if len(orders) == 0 {
		return nil
	}
	now := time.Now()
	if orders != nil {
		orderedAt := orders[0].UpdatedAt
		diff := now.Sub(orderedAt).Hours() / 24
		if diff > 14 {
			return nil
		}
	}
	return orders
}

func (s *ServiceOrder) CancelOrders(orders []Order) error {
	for _, o := range orders {
		o.IsOrder = false
		err := s.r.Update(o)
		if err != nil {
			return err
		}
		err = s.r.Delete(o)
		if err != nil {
			return err
		}
	}
	return nil
}
