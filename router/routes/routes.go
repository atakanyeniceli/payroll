package routes

import (
	"net/http"

	hourlyrate "github.com/atakanyeniceli/payroll/models/hourlyrate/handler"
	overtimeHandler "github.com/atakanyeniceli/payroll/models/overtime/handler"
	userHandler "github.com/atakanyeniceli/payroll/models/user/handler"
	webHandler "github.com/atakanyeniceli/payroll/web/handler"
)

// ----------------------------------------------------------------------
// 1. PUBLIC ROTALAR (Herkes Erişebilir)
// ----------------------------------------------------------------------

func PublicWebRoutes(mux *http.ServeMux, h *webHandler.Handler, fs http.Handler) {
	// Statik Dosyalar
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	// Sayfalar
	mux.HandleFunc("GET /login", h.Login)
	mux.HandleFunc("GET /register", h.Register)
	mux.HandleFunc("GET /{$}", h.Index)
}

func PublicAuthRoutes(mux *http.ServeMux, h *userHandler.Handler) {
	mux.HandleFunc("POST /login", h.WebLogin)
	mux.HandleFunc("POST /register", h.WebRegister)
	mux.HandleFunc("GET /logout", h.WebLogout)
}

// ----------------------------------------------------------------------
// 2. DASHBOARD ROTALARI (Korumalı Web Alanı - HTML/HTMX)
// ----------------------------------------------------------------------
// Not: Bu rotalar main.go'da "/dashboard" altına mount edilir.
// Buradaki "/" aslında "/dashboard/" demektir.

func DashboardUserRoutes(mux *http.ServeMux, h *userHandler.Handler) {
	mux.HandleFunc("GET /{$}", h.WebDashboard) // Dashboard Ana Sayfa
}

func DashboardOvertimeRoutes(mux *http.ServeMux, h *overtimeHandler.Handler) {
	// Modallar
	mux.HandleFunc("GET /modals/overtime", h.GetModal) // /dashboard/modals/overtime

	mux.HandleFunc("GET /overtime", h.GetDashboard)
	mux.HandleFunc("POST /overtime", h.Create) // /dashboard/overtime

	// İşlemler (HTMX Post)
	mux.HandleFunc("POST /save/overtime", h.Create) // /dashboard/save/overtime
}
func DashboardHourlyRateRoutes(mux *http.ServeMux, h *hourlyrate.Handler) {
	mux.HandleFunc("POST /hourlyrate", h.Create) // /dashboard/hourlyrate

}

// ----------------------------------------------------------------------
// 3. API ROTALARI (Mobil Uygulama - JSON)
// ----------------------------------------------------------------------
// Not: Bu rotalar main.go'da "/api" altına mount edilir.

func ApiUserRoutes(mux *http.ServeMux, h *userHandler.Handler) {
	// mux.HandleFunc("POST /login", h.ApiLogin) // Örnek: Mobil Login
}

func ApiOvertimeRoutes(mux *http.ServeMux, h *overtimeHandler.Handler) {
	// mux.HandleFunc("POST /overtime", h.CreateApi) // Örnek: Mobil Mesai Kayıt
}
