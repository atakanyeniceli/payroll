package service

import (
	"database/sql"
	"errors"
	"strings"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
	"github.com/atakanyeniceli/payroll/models/user"
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
		return user.ErrMissingFields
	}

	if !emailcheck.EmailCheck(email) {
		return user.ErrInvalidEmail
	}

	if password != confirmPassword {
		return user.ErrPasswordsNotMatch
	}

	// 3. Kullanıcının Mevcut Olup Olmadığını Kontrol Et
	existingUser, err := s.Repo.GetUserByEmail(email)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return apperror.NewServerError(err)
	}
	if existingUser != nil { // err == nil ise, kullanıcı bulunmuştur.
		return user.ErrEmailTaken
	}

	// Şifreyi Hash'le
	hashedPass, err := hash.Hash(password)
	if err != nil {
		return apperror.NewServerError(err)
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
		return apperror.NewServerError(err)
	}

	return nil
}
