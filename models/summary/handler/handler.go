package handler

import (
	"html/template"

	summaryService "github.com/atakanyeniceli/payroll/models/summary/service"
	"github.com/atakanyeniceli/payroll/tools/token"
)

type Handler struct {
	Tmpl         *template.Template
	Service      *summaryService.Service
	TokenManager *token.Manager
}

func NewHandler(s *summaryService.Service, t *template.Template, tm *token.Manager) *Handler {
	return &Handler{
		Service:      s,
		Tmpl:         t,
		TokenManager: tm,
	}
}
