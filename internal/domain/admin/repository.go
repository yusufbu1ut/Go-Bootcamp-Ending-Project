package admin

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type RepositoryAdmin struct {
	db *gorm.DB
}

func NewRepositoryAdmin(db *gorm.DB) *RepositoryAdmin {
	return &RepositoryAdmin{
		db: db,
	}
}

func (r *RepositoryAdmin) Migration() {
	r.db.AutoMigrate(&Admin{})
}

func (r *RepositoryAdmin) InsertSampleData() {
	admins := []Admin{{Name: "yusuf", AdminName: "yb", Email: "y@b.com", Password: "yusuf123"}}
	for _, a := range admins {
		r.db.Where(Admin{Email: a.Email}).Attrs(Admin{Email: a.Email, AdminName: a.AdminName}).FirstOrCreate(&a)
	}
}

func (r *RepositoryAdmin) GetAll(pageIndex, pageSize int) ([]Admin, int) {
	var admins []Admin
	var count int64

	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&admins).Count(&count)

	return admins, int(count)
}

func (r *RepositoryAdmin) GetByID(id int) Admin {
	var admin Admin
	result := r.db.Find(&admin, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Admin not found with id : %d", id)
		return Admin{}
	}
	return admin
}

func (r *RepositoryAdmin) GetByName(name string) []Admin {
	var admins []Admin
	r.db.Where("Name LIKE ?", "%"+name+"%").Find(&admins)
	return admins
}

func (r *RepositoryAdmin) GetByMail(mail string) []Admin {
	var admins []Admin
	r.db.Where("Email = ?", mail).Find(&admins)
	return admins
}

func (r *RepositoryAdmin) GetByMailAndPassword(mail string, pass string) []Admin {
	var admins []Admin
	r.db.Raw("SELECT * FROM Admin WHERE Email = ? AND Password =?", mail, pass).Scan(&admins)
	return admins
}

func (r *RepositoryAdmin) Delete(a Admin) error {
	result := r.db.Delete(a)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryAdmin) DeleteByID(id int) error {
	result := r.db.Delete(&Admin{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryAdmin) Create(a *Admin) error {
	result := r.db.Create(a)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryAdmin) Update(a Admin) error {
	result := r.db.Save(a)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
