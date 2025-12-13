package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func Init() *sql.DB {
	// 1. .env Dosyasını Yükle
	// godotenv.Load() ile dosyadaki değişkenler os.Getenv için hazır hale gelir.
	if err := godotenv.Load(); err != nil {
		// Hata olsa bile devam ederiz, çünkü değişkenler OS'tan gelmiş olabilir.
		log.Fatal("UYARI: .env dosyası bulunamadı. Ortam değişkenleri kullanılacak.")
	}

	// 2. Bağlantı Parametrelerini os.Getenv ile Oku
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbSSL := os.Getenv("DB_SSLMODE")

	// 3. Kontrol: Kritik değişkenler eksik mi?
	if dbUser == "" || dbName == "" {
		log.Fatal("kritik veritabanı değişkenleri (DB_USER/DB_NAME) eksik")
	}

	// 4. Bağlantı Dizesini (DSN) Oluştur
	dbURL := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPass, dbName, dbSSL,
	)

	// 5. Bağlantıyı Kur (sql.Open)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("sql.Open hatası: %v", err)
	}

	// 6. Bağlantı Havuzu Ayarları (Performans için önerilir)
	db.SetMaxOpenConns(25)                 // Maksimum açık bağlantı sayısı
	db.SetMaxIdleConns(25)                 // Boşta bekleyen bağlantı sayısı
	db.SetConnMaxLifetime(5 * time.Minute) // Bir bağlantının maksimum ömrü

	// 7. Bağlantıyı Test Et (Ping)
	if err = db.Ping(); err != nil {
		db.Close()
		log.Fatalf("veritabanı bağlantısı başarısız (ping): %v", err)
	}

	log.Println("✅ PostgreSQL bağlantısı başarıyla kuruldu.")
	return db
}
