package models

type Buku struct {
	ID          int    `json:"id"`
	IDBuku      int    `json:"id_buku"`
	Judul       string `json:"judul"`
	Pengarang   string `json:"pengarang"`
	Penulis     string `json:"penulis"`
	Penerbit    string `json:"penerbit"`
	TahunTerbit int    `json:"tahun_terbit"`
	Deskripsi   string `json:"deskripsi"`
	Gambar      string `json:"gambar"`
	Stok        int    `json:"stok"`
}
