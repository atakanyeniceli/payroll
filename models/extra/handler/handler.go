package handler

import (
	"html/template"

	"github.com/atakanyeniceli/payroll/models/extra/service"
	"github.com/atakanyeniceli/payroll/tools/token"
)

type Handler struct {
	Service *service.Service
	Tmpl    *template.Template
	Tm      *token.Manager
}

func NewHandler(s *service.Service, tmpl *template.Template, tm *token.Manager) *Handler {
	return &Handler{
		Service: s,
		Tmpl:    tmpl,
		Tm:      tm,
	}
}
