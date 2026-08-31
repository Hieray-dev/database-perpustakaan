package controllers

import (
	"encoding/json"
	"net/http"
	"perpustakaan-api/config"
)

type Penerbit struct {
	IDPenerbit   int    `json:"id_penerbit"`
	NamaPenerbit string `json:"nama_penerbit"`
}

func GetPenerbitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	rows, err := config.DB.Query("SELECT id_penerbit, nama_penerbit FROM penerbit")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var listPenerbit []Penerbit
	for rows.Next() {
		var p Penerbit
		if err := rows.Scan(&p.IDPenerbit, &p.NamaPenerbit); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		listPenerbit = append(listPenerbit, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listPenerbit)
}

func CreatePenerbitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var p Penerbit
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Format JSON tidak valid", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO penerbit (nama_penerbit) VALUES (?)"
	_, err = config.DB.Exec(query, p.NamaPenerbit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Penerbit berhasil ditambahkan!",
	})
}
