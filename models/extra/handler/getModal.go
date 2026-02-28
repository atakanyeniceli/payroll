package handler

import (
	"net/http"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
)

// GetModal: Prim ekleme penceresini (HTML) döndürür
func (h *Handler) GetModal(w http.ResponseWriter, r *http.Request) {
	// Template adının dosya adıyla (define block yoksa) aynı olduğundan emin olun
	// "extraAddModal.html" dosyasını template.Init() ile yüklediğinizi varsayıyoruz
	if err := h.Tmpl.ExecuteTemplate(w, "extraAddModal.html", nil); err != nil {
		code, msg := apperror.Resolve(err)
		http.Error(w, msg, code)
	}
}
