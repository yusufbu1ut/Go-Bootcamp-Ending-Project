package admin

import (
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/hashing"
)

type ServiceAdmin struct {
	r *RepositoryAdmin
}

func NewServiceAdmin(r *RepositoryAdmin) *ServiceAdmin {
	return &ServiceAdmin{
		r: r,
	}
}

func (s *ServiceAdmin) GetUser(email string, password string) *Admin {
	admin := s.r.GetByMail(email)
	passCheck := hashing.CheckWordHash(password, admin.Password)
	if admin.ID != 0 && passCheck {
		return admin
	}
	return &Admin{}
}

func (s *ServiceAdmin) GetUserWithId(id int) *Admin {
	admin := s.r.GetByID(id)
	if admin.ID != 0 {
		return admin
	}
	return &Admin{}
}
