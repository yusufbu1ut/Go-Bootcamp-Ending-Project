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
	admin1 := NewAdmin("yusufb", "y@b.com", "yusuf123")
	admin2 := NewAdmin("ysf", "y@sf.com", "yusuf123")
	admins := []Admin{
		*admin1,
		*admin2,
	}
	for _, a := range admins {

		r.db.Where(Admin{Email: a.Email}).Attrs(Admin{Email: a.Email, Username: a.Username}).FirstOrCreate(&a)
	}
}

func (r *RepositoryAdmin) GetAll(pageIndex, pageSize int) ([]Admin, int) {
	var admins []Admin
	var count int64

	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&admins)
	r.db.Model(&Admin{}).Count(&count)
	return admins, int(count)
}

func (r *RepositoryAdmin) GetByID(id int) *Admin {
	var admin Admin
	result := r.db.Find(&admin, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Admin not found with id : %d", id)
		return &Admin{}
	}
	return &admin
}

func (r *RepositoryAdmin) GetByUserName(name string) *Admin {
	var admin Admin
	result := r.db.Where("Username = ?", name).Find(&admin)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Admin not found with name : %s", name)
		return &Admin{}
	}
	return &admin
}

func (r *RepositoryAdmin) GetByMail(mail string) *Admin {
	var admin Admin
	result := r.db.Where("Email = ?", mail).First(&admin)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Admin not found with mail : %s", mail)
		return &Admin{}
	}
	return &admin
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
