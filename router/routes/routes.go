package routes

import (
	"net/http"

	userHandler "github.com/atakanyeniceli/payroll/models/user/handler"
	"github.com/atakanyeniceli/payroll/router"
	webHandler "github.com/atakanyeniceli/payroll/web/handler"
)

// RegisterPublicRoutes, herkese açık web sayfalarını ve statik dosyaları tanımlar.
func PublicWebRoutes(r *router.Router, h *webHandler.Handler, fs http.Handler) {
	r.Handle("GET /static/", http.StripPrefix("/static/", fs))
	r.HandleFunc("GET /login", h.Login)
	r.HandleFunc("GET /register", h.Register)
	r.HandleFunc("GET /{$}", h.Index)
}

func UserWebRoutes(r *router.Router, h *userHandler.Handler, authMiddleware func(http.Handler) http.Handler) {
	r.HandleFunc("POST /login", h.WebLogin)
	r.HandleFunc("POST /register", h.WebRegister)
	r.HandleFunc("GET /logout", h.WebLogout)
	r.Handle("GET /dashboard", authMiddleware(http.HandlerFunc(h.WebDashboard)))

}
