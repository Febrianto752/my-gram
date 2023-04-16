package helper

import "golang.org/x/crypto/bcrypt"

func HashPass(password string) string {
	salt := 8
	hashingPassword, _ := bcrypt.GenerateFromPassword([]byte(password), salt)

	return string(hashingPassword)
}

func ComparePass(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
