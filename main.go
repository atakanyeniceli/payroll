package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/atakanyeniceli/payroll/database"

	// Modüller
	extraHandler "github.com/atakanyeniceli/payroll/models/extra/handler"
	extraRepository "github.com/atakanyeniceli/payroll/models/extra/repository"
	extraService "github.com/atakanyeniceli/payroll/models/extra/service"

	summaryHandler "github.com/atakanyeniceli/payroll/models/summary/handler"
	summaryService "github.com/atakanyeniceli/payroll/models/summary/service"

	hourlyRateHandler "github.com/atakanyeniceli/payroll/models/hourlyrate/handler"
	hourlyRateRepository "github.com/atakanyeniceli/payroll/models/hourlyrate/repository"
	hourlyRateService "github.com/atakanyeniceli/payroll/models/hourlyrate/service"

	overtimeHandler "github.com/atakanyeniceli/payroll/models/overtime/handler"
	overtimeRepository "github.com/atakanyeniceli/payroll/models/overtime/repository"
	overtimeService "github.com/atakanyeniceli/payroll/models/overtime/service"

	userHandler "github.com/atakanyeniceli/payroll/models/user/handler"
	userRepository "github.com/atakanyeniceli/payroll/models/user/repository"
	userService "github.com/atakanyeniceli/payroll/models/user/service"

	"github.com/atakanyeniceli/payroll/router"
	"github.com/atakanyeniceli/payroll/router/routes"
	"github.com/atakanyeniceli/payroll/tools/logger"
	"github.com/atakanyeniceli/payroll/tools/token"
	webHandler "github.com/atakanyeniceli/payroll/web/handler"
	webTmpl "github.com/atakanyeniceli/payroll/web/template"
)

// Handlers struct'ı, init fonksiyonundan main'e temiz veri taşımak için kullanılır
type Handlers struct {
	User       *userHandler.Handler
	Overtime   *overtimeHandler.Handler
	HourlyRate *hourlyRateHandler.Handler
	Summary    *summaryHandler.Handler
	Extra      *extraHandler.Handler
	Web        *webHandler.Handler
}

func main() {
	// --- BAŞLATMA ---
	tmpl := webTmpl.Init()
	db := database.Init()
	logger.InitLogger()
	defer logger.LogFile.Close()
	defer db.Close()

	tokenManager := token.NewManager()

	// Handler'ları Hazırla
	h := handlerInit(db, tmpl, tokenManager)
	webFS := http.FileServer(http.Dir(filepath.Join(getWD(), "web")))

	// --- ROUTER KURULUMU ---
	r := router.NewRouter() // Ana Router

	// 1. PUBLIC ALAN
	// Login, Register ve Statik dosyalar
	routes.PublicRoutes(r.Mux, h.Web, h.User, webFS)

	// 2. PRIVATE WEB ALANI (Sub-Router)
	// Sadece yetkili kullanıcıların erişebileceği alan
	privateWebMux := http.NewServeMux()

	// Tüm özel web rotalarını tek seferde kaydet
	routes.PrivateWebRoutes(
		privateWebMux,
		h.User,
		h.Overtime,
		h.HourlyRate,
		h.Summary,
		h.Extra,
	)

	// Dashboard Middleware Bağlantısı
	webAuth := router.WebAuthMiddleware(tokenManager)
	r.Handle("/dashboard/", http.StripPrefix("/dashboard", webAuth(privateWebMux)))

	// --- SUNUCUYU BAŞLAT ---
	r.Run()
}

func handlerInit(db *sql.DB, tmpl *template.Template, tm *token.Manager) *Handlers {
	// 1. User
	userRepo := userRepository.NewRepository(db)
	userSrv := userService.NewService(userRepo)
	userH := userHandler.NewHandler(userSrv, tmpl, tm)

	// 2. Hourly Rate (Overtime buna bağımlı olduğu için önce başlattık)
	rateRepo := hourlyRateRepository.NewRepository(db)
	rateSrv := hourlyRateService.NewService(rateRepo)
	rateH := hourlyRateHandler.NewHandler(rateSrv, tmpl, tm)

	// 3. Overtime
	overtimeRepo := overtimeRepository.NewRepository(db)
	// Overtime servisi, maaş hesaplaması için Rate servisine ihtiyaç duyar
	overtimeSrv := overtimeService.NewService(overtimeRepo, rateSrv)
	overtimeH := overtimeHandler.NewHandler(overtimeSrv, tmpl, tm)

	//Summary
	summarySrv := summaryService.NewService(overtimeSrv, rateSrv)
	summaryH := summaryHandler.NewHandler(summarySrv, tmpl, tm)

	// --- EXTRA MODULE ---
	extraRepo := extraRepository.NewRepository(db)
	extraSrv := extraService.NewService(extraRepo)
	extraH := extraHandler.NewHandler(extraSrv, tmpl, tm)

	// 4. Web Pages
	webH := webHandler.NewHandler(tmpl)

	return &Handlers{
		User:       userH,
		Overtime:   overtimeH,
		HourlyRate: rateH,
		Summary:    summaryH,
		Extra:      extraH,
		Web:        webH,
	}
}

func getWD() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return wd
}
