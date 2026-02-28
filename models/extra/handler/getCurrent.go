package handler

import (
	"net/http"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
	"github.com/atakanyeniceli/payroll/router"
)

func (h *Handler) GetCurrent(w http.ResponseWriter, r *http.Request) {
	// Context'ten UserID'yi al (Middleware tarafından eklendi)
	userID, ok := r.Context().Value(router.UserIDKey).(int)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, err := h.Service.GetCurrent(r.Context(), userID)
	if err != nil {
		code, msg := apperror.Resolve(err)
		w.WriteHeader(code)
		w.Write([]byte(msg))
		return
	}

}
