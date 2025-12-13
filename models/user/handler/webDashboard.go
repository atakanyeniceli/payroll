package handler

import (
	"net/http"

	"github.com/atakanyeniceli/payroll/router"
)

func (h *Handler) WebDashboard(w http.ResponseWriter, r *http.Request) {
	// Middleware sayesinde bu handler'a ulaşıldığında kullanıcının kimliği doğrulanmıştır.
	// Request context'inden kullanıcı ID'sini alalım.
	_, ok := r.Context().Value(router.UserIDKey).(int)
	if !ok {
		// Bu durumun normalde olmaması gerekir, çünkü middleware kontrol ediyor.
		// Bir sorun varsa loglayıp hata dönmek en iyisidir.
		http.Error(w, "Kullanıcı bilgisi alınamadı.", http.StatusInternalServerError)
		return
	}

	// TODO: Bu userID ile servisi çağırıp kullanıcıya özel verileri çek ve template'i render et.
	err := h.Tmpl.ExecuteTemplate(w, "base.html", "dashboard")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
