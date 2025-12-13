package emailcheck

import "regexp"

// EmailCheck, verilen string'in geçerli bir e-posta formatında olup olmadığını kontrol eder.
// RFC 5322 standardına %100 uyumlu olmasa da çoğu durum için yeterlidir.
func EmailCheck(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
