package logger

import (
	"log"
	"os"
)

var LogFile *os.File // Dosyayı main'de kapatabilmek için dışarı açıyoruz

func InitLogger() {
	var err error

	// 1. Dosyayı Aç (veya Yoksa Oluştur)
	// O_APPEND: Dosyanın sonuna ekle (Eski logları silme)
	// O_CREATE: Dosya yoksa oluştur
	// O_WRONLY: Sadece yazma modunda aç
	// 0666: Okuma/Yazma izinleri
	LogFile, err = os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Log dosyası açılamadı: ", err)
	}

	// 2. Standart Log Çıktısını Dosyaya Yönlendir
	log.SetOutput(LogFile)

	// 3. Log Formatını Ayarla
	// Ldate: Tarih (2025/12/13)
	// Ltime: Saat (14:05:01)
	// Lshortfile: Hatanın olduğu dosya ve satır numarası (handler.go:45)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
