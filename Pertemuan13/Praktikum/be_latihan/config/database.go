package config

import (
	"log"
	"os"
	"fmt"

	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("../.env")
		if err != nil {
			log.Println(".env tidak ditemukan")
		}
	}

	dsn := os.Getenv("SUPABASE_DSN")
	if dsn == "" {
		log.Fatal("SUPABASE_DSN tidak ditemukan dalam variabel lingkungan")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal menghubungkan ke database: %v", err)
	}
	DB = db
	fmt.Println("Koneksi ke database berhasil")
}	

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database belum diinisialisasi. Pastikan InitDB() telah dipanggil.")
	}	
	return DB
}
