package handler

import (
	"log"
	"net/http"
	"strings"
)

func (h *Handler) WebRegister(w http.ResponseWriter, r *http.Request) {
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	err := h.Service.Register(firstname, lastname, email, password, confirmPassword)

	if err != nil {
		// 1. Loglama: Hatanın teknik detayını ve bağlamını (hangi email?) sunucu konsoluna yazıyoruz.
		log.Printf("Register Error [Email: %s]: %v", email, err)

		// 2. Hata Filtreleme: Teknik hataları (SQL, DB bağlantısı vb.) kullanıcıdan gizliyoruz.
		errMsg := err.Error()
		if strings.Contains(strings.ToLower(errMsg), "sql") || strings.Contains(strings.ToLower(errMsg), "connection refused") {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Sistemde geçici bir hata oluştu. Lütfen daha sonra tekrar deneyiniz."))
			return
		}

		// 3. Kullanıcı Hatası: Validasyon hatalarını olduğu gibi gösteriyoruz.
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errMsg))
		return
	}

	// Başarılı kayıt sonrası kullanıcıyı bilgilendirip login sayfasına yönlendirelim.
	// HTMX'in tetiklemesi için HX-Redirect header'ı kullanmak daha standart bir yöntemdir.
	w.Header().Set("HX-Redirect", "/login?status=success")
}
