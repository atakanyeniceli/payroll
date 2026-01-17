package handler

import (
	"html/template"

	"github.com/atakanyeniceli/payroll/models/hourlyrate/service"
	"github.com/atakanyeniceli/payroll/tools/token"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service, t *template.Template, tm *token.Manager) *Handler {
	return &Handler{Service: service}
}
