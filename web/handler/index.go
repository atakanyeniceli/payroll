package handler

import (
	"net/http"

	"github.com/atakanyeniceli/payroll/web/template"
)

func IndexHTML(w http.ResponseWriter, r *http.Request) {

	urlPath := r.URL.Path
	if urlPath != "/" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err := template.Tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

}
