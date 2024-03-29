package admin

import (
	"fmt"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/hashing"
	"gorm.io/gorm"
	"log"
)

type Admin struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAdmin(admin string, email string, pass string) *Admin {
	password, err := hashing.HashWord(pass)
	if err != nil {
		log.Println("Error occurred: ", err.Error())
	}
	return &Admin{
		Username: admin,
		Email:    email,
		Password: password,
	}
}

func (a *Admin) ToString() string {
	return fmt.Sprintf("Id: %d, Admin:%s, Mail: %s", a.ID, a.Username, a.Email)
}

func (a *Admin) BeforeDelete(tx *gorm.DB) (err error) {
	log.Printf("Admin %s deleting...", a.Username)
	return nil
}
