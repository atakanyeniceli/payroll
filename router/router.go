package router

import (
	"log"
	"net/http"
	"time"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) Handle(pattern string, handler http.Handler) {
	r.mux.Handle(pattern, handler)
}

func (r *Router) HandleFunc(pattern string, handlerFunc http.HandlerFunc) {
	r.mux.HandleFunc(pattern, handlerFunc)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

func (r *Router) Run() {

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("Sunucu http://localhost:8080 adresinde başlatılıyor...")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Sunucu başlatılamadı: %v", err)
	}
}
