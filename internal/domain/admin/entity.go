package admin

import (
	"fmt"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name      string `json:"name"`
	AdminName string `json:"admin-name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	LogStatus bool   `json:"status"`
}

func NewAdmin(name string, admin string, email string, pass string) *Admin {
	return &Admin{
		Name:      name,
		AdminName: admin,
		Email:     email,
		Password:  pass,
	}
}

func (a *Admin) ToString() string {
	return fmt.Sprintf("Id: %d, Name: %s, Admin:%s, Mail: %s", a.ID, a.Name, a.AdminName, a.Email)
}

func (a *Admin) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Printf("Admin %s deleting...", a.AdminName)
	return nil
}
