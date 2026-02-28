package routes

import (
	"net/http"

	extraHandler "github.com/atakanyeniceli/payroll/models/extra/handler"
	hourlyrateHandler "github.com/atakanyeniceli/payroll/models/hourlyrate/handler"
	overtimeHandler "github.com/atakanyeniceli/payroll/models/overtime/handler"
	summaryHandler "github.com/atakanyeniceli/payroll/models/summary/handler"
	userHandler "github.com/atakanyeniceli/payroll/models/user/handler"
	webHandler "github.com/atakanyeniceli/payroll/web/handler"
)

// ==============================================================================
// 1. PUBLIC ROUTES (Giriş Yapılmadan Erişilenler)
// ==============================================================================
// URL Prefix: / (Root)

func PublicRoutes(mux *http.ServeMux, wh *webHandler.Handler, uh *userHandler.Handler, fs http.Handler) {
	// Statik Dosyalar (CSS, JS, Images)
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	// Sayfalar (HTML Render)
	mux.HandleFunc("GET /login", wh.Login)
	mux.HandleFunc("GET /register", wh.Register)
	mux.HandleFunc("GET /{$}", wh.Index) // Landing Page

	// Auth İşlemleri (Form Post)
	mux.HandleFunc("POST /login", uh.WebLogin)
	mux.HandleFunc("POST /register", uh.WebRegister)
	mux.HandleFunc("GET /logout", uh.WebLogout)
}

// ==============================================================================
// 2. PRIVATE WEB ROUTES (Giriş Yapmış Kullanıcılar)
// ==============================================================================
// URL Prefix: /dashboard/ (Main.go'da strip edilir)
// Not: Buradaki yolların başına "/dashboard" yazmanıza gerek yoktur.

func PrivateWebRoutes(
	mux *http.ServeMux,
	uh *userHandler.Handler,
	oh *overtimeHandler.Handler,
	hh *hourlyrateHandler.Handler,
	sh *summaryHandler.Handler,
	eh *extraHandler.Handler,
) {
	// --- ANA SAYFA ---
	mux.HandleFunc("GET /{$}", uh.WebDashboard) // /dashboard

	// --- OVERTIME (Fazla Mesai) ---
	// Sayfa Parçaları (HTMX Partial)
	mux.HandleFunc("GET /overtime", oh.GetDashboard)    // Tabloyu getirir
	mux.HandleFunc("GET /modals/overtime", oh.GetModal) // Modalı getirir

	// İşlemler (Actions)
	mux.HandleFunc("POST /overtime", oh.Create)

	// --- HOURLY RATE (Saatlik Ücret) ---
	// Sayfa Parçaları
	mux.HandleFunc("GET /modals/hourlyrate", hh.GetModal)

	// İşlemler
	mux.HandleFunc("POST /hourlyrate", hh.Create)
	mux.HandleFunc("GET /hourlyrate", hh.GetCurrent)

	//----SUMMARY-----
	mux.HandleFunc("GET /summary", sh.GetCurrent)

	//------EXTRA(İkramiye,prim vb)---------
	mux.HandleFunc("GET /modals/extra", eh.GetModal)
	mux.HandleFunc("GET /extra", eh.GetCurrent)
	mux.HandleFunc("POST /extra", eh.Create)

}
