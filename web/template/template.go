package template

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Init() *template.Template {
	// 1. Çalışma dizinini al
	wdPath, err := os.Getwd()
	if err != nil {
		log.Fatal("Web Template init error=", err)
	}

	// 2. HTML dosyalarının olduğu kök klasörü belirle
	rootDir := filepath.Join(wdPath, "web", "html")

	// Formatlayıcı oluştur (İngilizce locale: 1,234.56 formatı için)
	// Eğer Türkçe format (1.234,56) isterseniz language.Turkish kullanabilirsiniz.
	p := message.NewPrinter(language.Turkish)

	// 3. Boş bir template oluştur
	tmpl := template.New("").Funcs(template.FuncMap{
		"formatMoney": func(n float64) string {
			return p.Sprintf("%.2f", n)
		},
	})

	// 4. filepath.Walk ile kök klasörden başlayarak tüm ağacı gez
	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		// Erişim hatası varsa raporla ama devam et (veya durdur)
		if err != nil {
			return err
		}

		// Eğer bu bir klasör değilse VE uzantısı .html ise
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			// Dosyayı mevcut template setine ekle (ParseFiles)
			_, err = tmpl.ParseFiles(path)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal("Web template walk error=", err)
	}

	return tmpl
}
