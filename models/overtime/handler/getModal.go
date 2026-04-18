package handler

import (
	"net/http"
	"strconv"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
	"github.com/atakanyeniceli/payroll/router"
)

func (h *Handler) GetModal(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(router.UserIDKey).(int)
	if !ok {
		http.Error(w, "Oturum hatası", http.StatusUnauthorized)
		return
	}

	// URL'de ID varsa "Düzenle (Edit)" Modudur
	idStr := r.PathValue("id")
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err == nil && id > 0 {
			overtimeRec, err := h.Service.GetByID(r.Context(), id, userID)
			if err != nil {
				code, msg := apperror.Resolve(err)
				http.Error(w, msg, code)
				return
			}
			// Datayı şablona gönder
			err = h.Tmpl.ExecuteTemplate(w, "overtimeModal.html", overtimeRec)
			if err != nil {
				code, msg := apperror.Resolve(err)
				http.Error(w, msg, code)
				return
			}
			return
		}
		return
	}

	// ID yoksa "Ekle (Create)" modudur (nil gönderiyoruz)
	if err := h.Tmpl.ExecuteTemplate(w, "overtimeModal.html", nil); err != nil {
		code, msg := apperror.Resolve(err)
		http.Error(w, msg, code)
	}
}
