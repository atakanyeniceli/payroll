package service

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/atakanyeniceli/payroll/models/user"
	customErrors "github.com/atakanyeniceli/payroll/tools/customerrors"
	"github.com/atakanyeniceli/payroll/tools/emailcheck"
	"github.com/atakanyeniceli/payroll/tools/hash"
)

func (s *Service) Register(firstname, lastname, email, password, confirmPassword string) error {

	// 1. Veri Normalleştirme ve Temizleme
	email = strings.ToLower(strings.TrimSpace(email))
	firstname = strings.TrimSpace(firstname)
	lastname = strings.TrimSpace(lastname)

	// 2. Temel Doğrulama (Validation)
	if firstname == "" || lastname == "" || email == "" {
		return customErrors.ErrMissingFields
	}

	if !emailcheck.EmailCheck(email) {
		return customErrors.ErrInvalidEmail
	}

	if password != confirmPassword {
		return customErrors.ErrInvalidPassword
	}

	// 3. Kullanıcının Mevcut Olup Olmadığını Kontrol Et
	existingUser, err := s.Repo.GetUserByEmail(email)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("kullanıcı kontrol edilirken bir veritabanı hatası oluştu: %w", err)
	}
	if existingUser != nil { // err == nil ise, kullanıcı bulunmuştur.
		return fmt.Errorf("Bu email adresi zaten kayıtlı")
	}

	// Şifreyi Hash'le
	hashedPass, err := hash.Hash(password)
	if err != nil {
		return fmt.Errorf("şifre hashlenirken bir hata oluştu: %w", err)
	}

	// Repository'ye yazma emrini ver

	newUser := &user.User{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Password:  hashedPass,
	}

	err = s.Repo.CreateUser(newUser)
	if err != nil {
		return fmt.Errorf("yeni kullanıcı oluşturulurken bir hata oluştu: %w", err)
	}

	return nil
}
