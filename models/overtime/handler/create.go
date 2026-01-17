package handler

import (
	"net/http"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
	"github.com/atakanyeniceli/payroll/models/overtime/service"
	"github.com/atakanyeniceli/payroll/router"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(router.UserIDKey).(int)
	if !ok {
		http.Error(w, "Oturum bilgisi bulunamadı.", http.StatusUnauthorized)
		return
	}

	// DTO Hazırla
	dto := service.CreateOvertimeDTO{
		UserID:      userID,
		DateStr:     r.FormValue("workDate"),
		StartTime:   r.FormValue("startTime"),
		EndTime:     r.FormValue("endTime"),
		Multiplier:  r.FormValue("multiplier"),
		Description: r.FormValue("description"),
	}
	// Servisi Çağır
	_, err := h.Service.CreateOvertime(ctx, dto)

	// Hata Yönetimi
	if err != nil {
		code, msg := apperror.Resolve(err) // ⭐ apperror.Resolve
		w.WriteHeader(code)
		w.Write([]byte(msg))
		return
	}

	// Başarılı - Sadece Overtime tablosunu yenile
	w.Header().Set("HX-Trigger", "overtimeUpdated")
	w.WriteHeader(http.StatusOK)
}
