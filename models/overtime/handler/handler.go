package handler

import (
	"html/template"

	overtimeService "github.com/atakanyeniceli/payroll/models/overtime/service"
	"github.com/atakanyeniceli/payroll/tools/token"
)

type Handler struct {
	Tmpl         *template.Template
	Service      *overtimeService.Service
	TokenManager *token.Manager
}

func NewHandler(s *overtimeService.Service, t *template.Template, tm *token.Manager) *Handler {
	return &Handler{
		Tmpl:         t,
		Service:      s,
		TokenManager: tm,
	}
}
