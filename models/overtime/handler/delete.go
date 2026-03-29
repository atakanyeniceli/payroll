package handler

import (
	"net/http"
	"strconv"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
	"github.com/atakanyeniceli/payroll/router"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	// 1. Context'ten Güvenli UserID'yi al
	userID, ok := r.Context().Value(router.UserIDKey).(int)
	if !ok {
		http.Error(w, "Oturum hatası: Kullanıcı ID bulunamadı.", http.StatusUnauthorized)
		return
	}

	// 2. URL'den ID değerini al (Örn: /api/delete/overtime/5 -> id=5)
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Geçersiz kayıt kimliği", http.StatusBadRequest)
		return
	}

	// 3. Servisi çağırarak silme işlemini yap
	err = h.Service.Delete(r.Context(), id, userID)
	if err != nil {
		code, msg := apperror.Resolve(err)
		http.Error(w, msg, code)
		return
	}

	// 4. İşlem başarılı. Tablonun ve maaş özetinin kendini yenilemesi için Trigger gönder
	w.Header().Set("HX-Trigger", "overtimeUpdated, summaryUpdated")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mesai kaydı başarıyla silindi"))
}
