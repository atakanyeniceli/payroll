package handler

import (
	"net/http"
	"strconv"
	"time"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
	"github.com/atakanyeniceli/payroll/router" // UserIDKey'e erişmek için gerekli
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	//1. router.UserIDKey kullanarak context'teki değeri alıyoruz ve int tipine dönüştürüyoruz (type assertion).
	userID, ok := r.Context().Value(router.UserIDKey).(int)
	if !ok {
		// Eğer context'te bu veri yoksa veya tipi yanlışsa (Middleware çalışmamışsa)
		http.Error(w, "Oturum hatası: Kullanıcı ID bulunamadı.", http.StatusUnauthorized)
		return
	}

	// 2. Form verilerini parse et
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Form verisi okunamadı", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	amountStr := r.FormValue("amount")
	dateStr := r.FormValue("date")

	// 3. Tip dönüşümleri
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "Geçersiz tutar formatı", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Geçersiz tarih formatı", http.StatusBadRequest)
		return
	}

	// 4. Servisi çağır
	err = h.Service.Create(r.Context(), userID, title, amount, date)
	if err != nil {
		code, msg := apperror.Resolve(err)
		http.Error(w, msg, code)
		return
	}

	// 5. Başarılı Yanıt
	w.Header().Set("HX-Trigger", "extraUpdated, summaryUpdated")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Prim başarıyla eklendi!"))
}
