package user

import (
	"net/http"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
)

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

var (
	ErrUserNotFound      = apperror.NewClientError("Kullanıcı bulunamadı", http.StatusNotFound)
	ErrEmailTaken        = apperror.NewClientError("Bu e-posta zaten kullanımda", http.StatusConflict)
	ErrInvalidPass       = apperror.NewClientError("Şifre hatalı", http.StatusUnauthorized)
	ErrMissingFields     = apperror.NewClientError("Lütfen tüm zorunlu alanları doldurunuz", http.StatusBadRequest)
	ErrInvalidEmail      = apperror.NewClientError("Geçersiz e-posta adresi", http.StatusBadRequest)
	ErrPasswordsNotMatch = apperror.NewClientError("Şifreler uyuşmuyor", http.StatusBadRequest)
)
