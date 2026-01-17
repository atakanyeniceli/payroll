package handler

import (
	"net/http"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
)

func (h *Handler) WebRegister(w http.ResponseWriter, r *http.Request) {
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	err := h.Service.Register(firstname, lastname, email, password, confirmPassword)
	if err != nil {
		code, msg := apperror.Resolve(err)
		w.WriteHeader(code)
		w.Write([]byte(msg))
		return
	}
	// Başarılı kayıt sonrası kullanıcıyı bilgilendirip login sayfasına yönlendirelim.
	// HTMX'in tetiklemesi için HX-Redirect header'ı kullanmak daha standart bir yöntemdir.
	w.Header().Set("HX-Redirect", "/")
}
