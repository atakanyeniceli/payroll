package service

import (
	"fmt"
	"strings"

	"github.com/atakanyeniceli/payroll/models/user"
	"github.com/atakanyeniceli/payroll/tools/hash"
)

func (s *Service) Login(email, password string) (*user.User, error) {
	// E-posta adresini standart bir formata getir
	email = strings.ToLower(strings.TrimSpace(email))

	// Repository'den ham veriyi çek
	dbUser, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		// Kullanıcı bulunamadı veya veritabanı hatası durumunda aynı genel hatayı dön.
		// Bu, kullanıcı enumerasyon saldırılarını önler.
		return nil, fmt.Errorf("e-posta veya şifre hatalı")
	}

	// Güvenlik: Şifreyi kontrol et
	if !hash.CheckHash(password, dbUser.Password) {
		return nil, fmt.Errorf("e-posta veya şifre hatalı")
	}

	return dbUser, nil // Başarılı
}
