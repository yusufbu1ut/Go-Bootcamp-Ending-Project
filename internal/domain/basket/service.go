package basket

import (
	"github.com/google/uuid"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/order"
)

type ServiceBasket struct {
	r *RepositoryBasket
}

func NewServiceBasket(r *RepositoryBasket) *ServiceBasket {
	return &ServiceBasket{
		r: r,
	}
}

func (s *ServiceBasket) GetAll(id int) []order.Order {
	baskets := s.r.GetByCustomerID(id)
	return baskets
}

func (s *ServiceBasket) AddBasket(baskets []order.Order) error {
	for _, basket := range baskets {
		basketItem := s.r.GetByIDsAndIsOrder(basket)
		if basketItem.ID == 0 || basketItem.IsOrder {
			err := s.r.Create(&basket)
			if err != nil {
				return err
			}
		} else {
			//checking is there same unordered item in db if there is only amount will be updated
			basketItem.Amount = basketItem.Amount + basket.Amount
			err := s.r.Update(basketItem)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *ServiceBasket) DeleteBasket(basket order.Order) error {
	err := s.r.Delete(basket)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceBasket) CompleteOrder(id int) error {
	basketItems := s.r.GetByCustomerID(id)
	//when order complete the order code generates for basket items
	orderCode := uuid.New().String()
	orderCode = orderCode[:8]
	for _, b := range basketItems {
		b.IsOrder = true
		b.OrderCode = orderCode
		err := s.r.Update(b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *ServiceBasket) GetBasket(basket order.Order) order.Order {
	basketItem := s.r.GetByIDsAndIsOrder(basket)
	if basketItem.ID == 0 {
		return order.Order{}
	}
	return basketItem
}

func (s *ServiceBasket) Update(basket order.Order) error {
	basketItem := s.r.GetByID(int(basket.ID))
	if basketItem.ID == 0 {
		return ErrItemNotExist
	}
	err := s.r.Update(basket)
	if err != nil {
		return err
	}
	return nil
}
