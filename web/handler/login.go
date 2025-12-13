package handler

import "net/http"

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Redirect", "/")
}
