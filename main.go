package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/atakanyeniceli/payroll/database"
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

// Handlers, uygulamanın tüm HTTP handler'larını tek bir çatı altında toplar.
// Bu yapı, handler sayısı arttıkça fonksiyon imzalarının şişmesini engeller.
type Handlers struct {
	User *userHandler.Handler
	Web  *webHandler.Handler
}

func main() {
	tmpl := webTmpl.Init()
	db := database.Init()
	logger.InitLogger()

	defer logger.LogFile.Close()
	defer db.Close()

	osWd := getWD()
	tokenManager := token.NewManager()

	h := handlerInit(db, tmpl, tokenManager)
	webFS := http.FileServer(http.Dir(filepath.Join(osWd, "web")))

	webAuth := router.WebAuthMiddleware(tokenManager)
	r := router.NewRouter()

	routes.PublicWebRoutes(r, h.Web, webFS)
	routes.UserWebRoutes(r, h.User, webAuth)

	r.Run()

}

func handlerInit(db *sql.DB, tmpl *template.Template, tm *token.Manager) *Handlers {
	userRepo := userRepository.NewRepository(db)
	userService := userService.NewService(userRepo)
	userHandler := userHandler.NewHandler(userService, tmpl, tm)

	webHandler := webHandler.NewHandler(tmpl)

	return &Handlers{
		User: userHandler,
		Web:  webHandler,
	}
}

func getWD() string {
	osWd, err := os.Getwd()
	if err != nil {
		log.Fatal("GETWD error=", err)
	}
	return osWd
}
