package controllers

import (
	"encoding/json"
	"net/http"
	"perpustakaan-api/config"
	"perpustakaan-api/models"
)

func GetKatalogBuku(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, `{"message":"Method tidak diizinkan"}`, http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	rows, err := config.DB.Query("SELECT id_buku, judul, COALESCE(penulis, ''), COALESCE(deskripsi, ''), COALESCE(gambar, ''), stok FROM buku")
	if err != nil {
		http.Error(w, `{"message":"Gagal mengambil data buku: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var listBuku []models.Buku
	for rows.Next() {
		var b models.Buku
		err := rows.Scan(&b.IDBuku, &b.Judul, &b.Penulis, &b.Deskripsi, &b.Gambar, &b.Stok)
		if err != nil {
			http.Error(w, `{"message":"Gagal membaca data buku"}`, http.StatusInternalServerError)
			return
		}
		listBuku = append(listBuku, b)
	}

	json.NewEncoder(w).Encode(listBuku)
}

func CreateBukuHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, `{"message":"Method tidak diizinkan"}`, http.StatusMethodNotAllowed)
		return
	}

	var b models.Buku
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, `{"message":"Format JSON tidak valid"}`, http.StatusBadRequest)
		return
	}

	query := `INSERT INTO buku (judul, penulis, deskripsi, gambar, stok) VALUES (?, ?, ?, ?, ?)`
	_, err = config.DB.Exec(query, b.Judul, b.Penulis, b.Deskripsi, b.Gambar, b.Stok)
	if err != nil {
		http.Error(w, `{"message":"Gagal menambah buku: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Buku berhasil ditambahkan!",
	})
}
