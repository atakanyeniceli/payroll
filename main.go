package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/atakanyeniceli/payroll/database"

	// User Modülü
	userHandler "github.com/atakanyeniceli/payroll/models/user/handler"
	userRepository "github.com/atakanyeniceli/payroll/models/user/repository"
	userService "github.com/atakanyeniceli/payroll/models/user/service"

	// Overtime Modülü
	overtimeHandler "github.com/atakanyeniceli/payroll/models/overtime/handler"
	overtimeRepository "github.com/atakanyeniceli/payroll/models/overtime/repository"
	overtimeService "github.com/atakanyeniceli/payroll/models/overtime/service"

	//Hourlyrate Modülü
	hourlyRateHandler "github.com/atakanyeniceli/payroll/models/hourlyrate/handler"
	hourlyRateRepository "github.com/atakanyeniceli/payroll/models/hourlyrate/repository"
	hourlyRateService "github.com/atakanyeniceli/payroll/models/hourlyrate/service"

	"github.com/atakanyeniceli/payroll/router"
	"github.com/atakanyeniceli/payroll/router/routes"
	"github.com/atakanyeniceli/payroll/tools/logger"
	"github.com/atakanyeniceli/payroll/tools/token"
	webHandler "github.com/atakanyeniceli/payroll/web/handler"
	webTmpl "github.com/atakanyeniceli/payroll/web/template"
)

// Handlers struct'ını genişlettik
type Handlers struct {
	User       *userHandler.Handler
	Overtime   *overtimeHandler.Handler
	HourlyRate *hourlyRateHandler.Handler
	Web        *webHandler.Handler
}

func main() {
	// 1. Temel Başlatmalar
	tmpl := webTmpl.Init()
	db := database.Init()
	logger.InitLogger()

	defer logger.LogFile.Close()
	defer db.Close()

	osWd := getWD()
	tokenManager := token.NewManager()

	// 2. Handler ve Servislerin Kurulumu
	h := handlerInit(db, tmpl, tokenManager)
	webFS := http.FileServer(http.Dir(filepath.Join(osWd, "web")))

	// 3. Router Kurulumu
	r := router.NewRouter() // Root Router

	// A) PUBLIC ROTALAR (Middleware Yok)
	// ------------------------------------------------------------
	// Login, Register, Static Files
	routes.PublicWebRoutes(r.Mux, h.Web, webFS)
	routes.PublicAuthRoutes(r.Mux, h.User)

	// B) DASHBOARD ROTALARI (Web Auth Middleware Var)
	// ------------------------------------------------------------
	// "/dashboard" ön eki ile gelen istekler buraya

	dashboardMux := http.NewServeMux()

	// Rotaları ekle
	routes.DashboardUserRoutes(dashboardMux, h.User)
	routes.DashboardOvertimeRoutes(dashboardMux, h.Overtime)
	routes.DashboardHourlyRateRoutes(dashboardMux, h.HourlyRate)

	// Middleware'i uygula ve Root Router'a bağla
	// StripPrefix: "/dashboard" kısmını url'den siler, iç router sadece "/modals/overtime" görür.
	webAuth := router.WebAuthMiddleware(tokenManager)
	r.Handle("/dashboard/", http.StripPrefix("/dashboard", webAuth(dashboardMux)))

	// C) API ROTALARI (Mobil - Token Auth Middleware Var)
	// ------------------------------------------------------------
	// Gelecekte burayı aktif edebilirsin
	/*
		apiMux := http.NewServeMux()
		routes.ApiOvertimeRoutes(apiMux, h.Overtime)
		apiAuth := router.ApiAuthMiddleware(tokenManager) // Farklı bir middleware olabilir
		r.Handle("/api/", http.StripPrefix("/api", apiAuth(apiMux)))
	*/

	// 4. Sunucuyu Başlat
	r.Run()
}

func handlerInit(db *sql.DB, tmpl *template.Template, tm *token.Manager) *Handlers {
	// User Modülü
	userRepo := userRepository.NewRepository(db)
	userService := userService.NewService(userRepo)
	userHandler := userHandler.NewHandler(userService, tmpl, tm)

	// Overtime Modülü (Eklendi)
	overtimeRepo := overtimeRepository.NewRepository(db)
	overtimeService := overtimeService.NewService(overtimeRepo)
	overtimeHandler := overtimeHandler.NewHandler(overtimeService, tmpl, tm) // Handler yapına göre parametreleri düzenle

	// HourlyRate Modülü
	hourlyRateRepo := hourlyRateRepository.NewRepository(db)
	hourlyRateService := hourlyRateService.NewService(hourlyRateRepo)
	hourlyRateHandler := hourlyRateHandler.NewHandler(hourlyRateService, tmpl, tm)

	// Web (Sayfa) Modülü
	webHandler := webHandler.NewHandler(tmpl)

	return &Handlers{
		User:       userHandler,
		Overtime:   overtimeHandler,
		HourlyRate: hourlyRateHandler,
		Web:        webHandler,
	}
}

func getWD() string {
	osWd, err := os.Getwd()
	if err != nil {
		log.Fatal("GETWD error=", err)
	}
	return osWd
}
