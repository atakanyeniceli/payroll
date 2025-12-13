package handler

import "html/template"

type Handler struct{ Tmpl *template.Template }

func NewHandler(tmpl *template.Template) *Handler {
	return &Handler{Tmpl: tmpl}
}
