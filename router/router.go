package router

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Router, standart http.ServeMux'u sarmalar.
type Router struct {
	Mux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		Mux: http.NewServeMux(),
	}
}

func (r *Router) Handle(pattern string, handler http.Handler) {
	r.Mux.Handle(pattern, handler)
}

func (r *Router) HandleFunc(pattern string, handlerFunc http.HandlerFunc) {
	r.Mux.HandleFunc(pattern, handlerFunc)
}

// ServeHTTP, Router'ın standart bir http.Handler gibi davranmasını sağlar.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.Mux.ServeHTTP(w, req)
}

func (r *Router) Run() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Println("🚀 Sunucu http://localhost:8080 adresinde başlatılıyor...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Sunucu başlatılamadı: %v", err)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Sunucu kapatılıyor...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Sunucu zorla kapatıldı: %v", err)
	}

	log.Println("✅ Sunucu başarıyla durduruldu.")
}
