package handler

import (
	"net/http"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
	"github.com/atakanyeniceli/payroll/router"
)

func (h *Handler) GetDashboard(w http.ResponseWriter, r *http.Request) {
	// Middleware üzerinden gelen UserID bilgisini al
	userID, ok := r.Context().Value(router.UserIDKey).(int)
	if !ok {
		http.Error(w, "Oturum bilgisi bulunamadı.", http.StatusUnauthorized)
		return
	}

	// Kullanıcıya ait verileri servis katmanından çek
	overtimes, err := h.Service.GetCurrentMonthLog(r.Context(), userID)
	if err != nil {
		code, msg := apperror.Resolve(err)
		http.Error(w, msg, code)
		return
	}

	// Veriyi template'e bas (overtimeList template'inin tanımlı olması gerekir)
	if err := h.Tmpl.ExecuteTemplate(w, "overtimeList.html", overtimes); err != nil {
		code, msg := apperror.Resolve(err)
		http.Error(w, msg, code)
	}
}
