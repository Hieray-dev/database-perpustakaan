package controllers

import (
	"encoding/json"
	"net/http"
	"perpustakaan-api/config"
)

type Genre struct {
	IDGenre   int    `json:"id_genre"`
	NamaGenre string `json:"nama_genre"`
}

func GetGenreHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	rows, err := config.DB.Query("SELECT id_genre, nama_genre FROM genre")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var listGenre []Genre
	for rows.Next() {
		var g Genre
		if err := rows.Scan(&g.IDGenre, &g.NamaGenre); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		listGenre = append(listGenre, g)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listGenre)
}

func CreateGenreHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var g Genre
	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		http.Error(w, "Format JSON tidak valid", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO genre (nama_genre) VALUES (?)"
	_, err = config.DB.Exec(query, g.NamaGenre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Genre berhasil ditambahkan!",
	})
}
