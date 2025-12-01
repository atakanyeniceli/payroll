package template

import (
	"html/template"
	"os"
	"path/filepath"
)

var Tmpl *template.Template

func Init() error {

	wdPath, err := os.Getwd()
	if err != nil {
		return nil
	}

	htmlFilePath := filepath.Join(wdPath, "web", "html", "*.html")

	Tmpl, err = template.ParseGlob(htmlFilePath)
	if err != nil {
		return err
	}
	return nil
}
