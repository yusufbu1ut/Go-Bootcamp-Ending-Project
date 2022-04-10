package hashing

import "golang.org/x/crypto/bcrypt"

//HashWord use to hash password and roles
func HashWord(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckWordHash checks hashed words is it same with input content
func CheckWordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
