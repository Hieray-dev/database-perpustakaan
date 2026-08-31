package controllers

import (
	"encoding/json"
	"net/http"
	"perpustakaan-api/config"
)

type Fasilitas struct {
	IDFasilitas   int    `json:"id_fasilitas"`
	NamaFasilitas string `json:"nama_fasilitas"`
	Kondisi       string `json:"kondisi"`
}

func GetFasilitasHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	rows, err := config.DB.Query("SELECT id_fasilitas, nama_fasilitas, COALESCE(kondisi, '') FROM fasilitas")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var listFasilitas []Fasilitas
	for rows.Next() {
		var f Fasilitas
		if err := rows.Scan(&f.IDFasilitas, &f.NamaFasilitas, &f.Kondisi); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		listFasilitas = append(listFasilitas, f)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listFasilitas)
}

func CreateFasilitasHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var f Fasilitas
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		http.Error(w, "Format JSON tidak valid", http.StatusBadRequest)
		return
	}

	if f.Kondisi == "" {
		f.Kondisi = "Baik"
	}

	query := "INSERT INTO fasilitas (nama_fasilitas, kondisi) VALUES (?, ?)"
	_, err = config.DB.Exec(query, f.NamaFasilitas, f.Kondisi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Fasilitas berhasil ditambahkan!",
	})
}
