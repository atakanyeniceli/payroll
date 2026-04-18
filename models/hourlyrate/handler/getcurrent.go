package handler

import (
	"net/http"
	"time"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
	"github.com/atakanyeniceli/payroll/router"
)

func (h *Handler) GetCurrent(w http.ResponseWriter, r *http.Request) {
	// 1. UserID al
	userID, ok := r.Context().Value(router.UserIDKey).(int)
	if !ok {
		http.Error(w, "Oturum hatası", http.StatusUnauthorized)
		return
	}

	// 2. Servisten Veriyi Çek (sadece amount verisi geliyor.)
	hr, err := h.Service.GetByIDAndDate(r.Context(), userID, time.Now())

	// Hata yönetimi (Kayıt yoksa rate nil gelebilir, bu normaldir)
	if err != nil {
		code, msg := apperror.Resolve(err)
		http.Error(w, msg, code)
		return
	}

	data := map[string]interface{}{
		"Amount":        hr.Amount,
		"EffectiveDate": hr.EffectiveDate,
		"ID":            hr.ID,
		"UserID":        userID,
	}

	// 3. Template Render Et
	if err := h.Tmpl.ExecuteTemplate(w, "hourlyrateContainer.html", data); err != nil {
		code, msg := apperror.Resolve(err)
		http.Error(w, msg, code)
	}
}
