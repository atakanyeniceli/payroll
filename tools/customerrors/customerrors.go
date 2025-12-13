package customErrors

import "errors"

var (
	ErrUserNotFound    = errors.New("Kullanıcı bulunamadı")
	ErrMissingFields   = errors.New("ad, soyad ve e-posta alanları zorunludur")
	ErrInvalidPassword = errors.New("Şifre hatalı")
	ErrInvalidEmail    = errors.New("geçersiz e-posta formatı")
	ErrEmailTaken      = errors.New("Bu e-posta adresi zaten kullanımda")
	ErrDatabase        = errors.New("Veritabanı hatası")
)
