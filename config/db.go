package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "root:ray@tcp(127.0.0.1:3306)/db_perpustakaan?parseTime=true"

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal konek database:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Database tidak merespon:", err)
	}

	fmt.Println("Koneksi Database Berhasil!")
}
