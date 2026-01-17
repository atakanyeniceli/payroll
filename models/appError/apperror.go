package apperror

import (
	"log"
	"net/http"
)

// AppError: Projenin ortak hata formatı
type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Log     error  `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

func Resolve(err error) (int, string) {
	if appErr, ok := err.(*AppError); ok {
		if appErr.Log != nil {
			log.Printf("SERVER ERROR | Code: %d | UserMsg: %s | Trace: %v", appErr.Code, appErr.Message, appErr.Log)
		}
		return appErr.Code, appErr.Message
	} else {
		// AppError değilse (örn: panic vb.) bunu da logla
		log.Printf("UNHANDLED ERROR: %v", err)
	}
	return http.StatusInternalServerError, "Sunucu tarafında teknik bir sorun oluştu"
}

// 400 Serisi Hata Üretici (Kullanıcı Hatası)
func NewClientError(msg string, code int) *AppError {
	return &AppError{
		Message: msg,
		Code:    code,
		Log:     nil,
	}
}

// 500 Serisi Hata Üretici (Sunucu Hatası)
func NewServerError(err error) *AppError {
	return &AppError{
		Message: "Sunucu tarafında teknik bir sorun oluştu",
		Code:    http.StatusInternalServerError,
		Log:     err,
	}
}
