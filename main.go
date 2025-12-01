package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/atakanyeniceli/payroll/router"
	webHandler "github.com/atakanyeniceli/payroll/web/handler"
	webTmpl "github.com/atakanyeniceli/payroll/web/template"
)

func main() {
	err := webTmpl.Init()
	if err != nil {
		log.Fatal("HTML init error =", err)
	}

	osWd, err := os.Getwd()
	if err != nil {
		log.Fatal("GETWD error=", err)
	}
	webFS := http.FileServer(http.Dir(filepath.Join(osWd, "web")))

	r := router.NewRouter()
	r.Handle("GET /static/", http.StripPrefix("/static/", webFS))
	r.HandleFunc("GET /", webHandler.IndexHTML)

	fmt.Print("Running...")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Router run error =", err)
	}
}
