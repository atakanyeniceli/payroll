package handler

import (
	"net/http"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
	"github.com/atakanyeniceli/payroll/router"
)

// GetCurrent: Ek ödemeler (Primler) tablosunu HTML olarak doldurur.
func (h *Handler) GetCurrent(w http.ResponseWriter, r *http.Request) {
	// 1. Context'ten UserID'yi güvenli şekilde al
	userID, ok := r.Context().Value(router.UserIDKey).(int)
	if !ok {
		http.Error(w, "Oturum hatası: Kullanıcı kimliği doğrulanamadı.", http.StatusUnauthorized)
		return
	}

	// 2. Servisten kullanıcının bu aya ait ekstra ödemelerini çek
	// (Not: Service metodunuzun adının GetCurrent olduğunu varsayıyorum)
	extras, err := h.Service.GetCurrent(r.Context(), userID)
	if err != nil {
		code, msg := apperror.Resolve(err)
		http.Error(w, msg, code)
		return
	}

	// 3. Veriyi oluşturduğumuz şablon ile render edip HTMX'e gönder
	if err := h.Tmpl.ExecuteTemplate(w, "extraList.html", extras); err != nil {
		code, msg := apperror.Resolve(err)
		http.Error(w, msg, code)
	}
}
