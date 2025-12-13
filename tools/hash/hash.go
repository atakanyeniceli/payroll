package hash

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (string, error) {
	hashedByes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedByes), nil
}

func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	// Hata yoksa (eşleşiyorsa) true döner.
	return err == nil
}
