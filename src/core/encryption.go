// core/encryption.go
package core

import "golang.org/x/crypto/bcrypt"

// EncryptPassword encripta una cadena usando bcrypt
func EncryptPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// VerifyPassword compara una cadena con su versión encriptada
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
