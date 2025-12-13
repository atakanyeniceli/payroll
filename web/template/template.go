package template

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func Init() *template.Template {

	wdPath, err := os.Getwd()
	if err != nil {
		log.Fatal("Web Template init error=", err)
	}

	htmlFilePath := filepath.Join(wdPath, "web", "html", "*.html")

	Tmpl, err := template.ParseGlob(htmlFilePath)
	if err != nil {
		log.Fatal("Web template init error=", err)
	}

	return Tmpl
}
