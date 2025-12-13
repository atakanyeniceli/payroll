package handler

import (
	"html/template"

	userService "github.com/atakanyeniceli/payroll/models/user/service"
	"github.com/atakanyeniceli/payroll/tools/token"
)

type Handler struct {
	Tmpl         *template.Template
	Service      *userService.Service
	TokenManager *token.Manager
}

func NewHandler(s *userService.Service, t *template.Template, tm *token.Manager) *Handler {
	return &Handler{
		Tmpl:         t,
		Service:      s,
		TokenManager: tm,
	}
}
