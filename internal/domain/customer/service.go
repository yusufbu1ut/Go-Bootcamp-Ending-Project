package customer

import (
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/hashing"
)

type ServiceCustomer struct {
	r *RepositoryCustomer
}

func NewServiceCustomer(r *RepositoryCustomer) *ServiceCustomer {
	return &ServiceCustomer{
		r: r,
	}
}

func (s *ServiceCustomer) Create(customer *Customer) error {
	if customer.Username == "" {
		return ErrCustomerUsernameNil
	}
	if customer.Email == "" {
		return ErrCustomerEmailNil
	}
	if customer.Password == "" {
		return ErrCustomerPasswordNil
	}
	exist := s.r.GetByUserName(customer.Username)
	if exist.ID != 0 {
		return ErrCustomerExistWithUsername
	}
	exist = s.r.GetByMail(customer.Email)
	if exist.ID != 0 {
		return ErrCustomerExistWithEmail
	}
	err := s.r.Create(customer)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceCustomer) GetUser(email string, password string) *Customer {
	user := s.r.GetByMail(email)
	passCheck := hashing.CheckWordHash(password, user.Password)
	if user.ID != 0 && passCheck {
		return user
	}
	return &Customer{}
}

func (s *ServiceCustomer) GetUserWithId(id int) *Customer {
	customer := s.r.GetByID(id)
	return customer
}
