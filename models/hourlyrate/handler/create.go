package handler

import (
	"net/http"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
	"github.com/atakanyeniceli/payroll/models/hourlyrate/service"
	"github.com/atakanyeniceli/payroll/router"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// 1. Context'ten UserID al
	userID, ok := ctx.Value(router.UserIDKey).(int)
	if !ok {
		http.Error(w, "Oturum hatası: Kullanıcı ID bulunamadı.", http.StatusUnauthorized)
		return
	}

	// 2. DTO Hazırla (Form verileriyle)
	dto := service.CreateHourlyRateDTO{
		UserID:           userID,
		AmountStr:        r.FormValue("rate"),
		EffectiveDateStr: r.FormValue("effectiveDate"),
	}

	// 3. Servisi Çağır
	_, err := h.Service.Create(ctx, dto)
	if err != nil {
		code, msg := apperror.Resolve(err)
		w.WriteHeader(code)
		w.Write([]byte(msg))
		return
	}

	// 4. Başarılı işlem sonrası yönlendirme
	w.Header().Set("HX-Trigger", "hourlyRateUpdated")
	w.WriteHeader(http.StatusOK)
}
