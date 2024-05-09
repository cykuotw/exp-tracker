package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ValidatePassword(hashed string, unhashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(unhashed))
	return err == nil
}
