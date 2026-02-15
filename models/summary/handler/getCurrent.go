package handler

import (
	"net/http"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
	"github.com/atakanyeniceli/payroll/router"
)

func (h *Handler) GetCurrent(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(router.UserIDKey).(int)
	if !ok {
		http.Error(w, "Oturum hatası", http.StatusUnauthorized)
		return
	}

	stats, err := h.Service.GetCurrent(r.Context(), userID)
	if err != nil {
		code, msg := apperror.Resolve(err)
		http.Error(w, msg, code)
		return
	}

	if err := h.Tmpl.ExecuteTemplate(w, "summary.html", stats); err != nil {
		code, msg := apperror.Resolve(err)
		http.Error(w, msg, code)
	}
}
