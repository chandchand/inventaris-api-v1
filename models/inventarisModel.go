package models

type Inventaris struct {
	ID                     int     `json:"id"`
	NamaBarang             string  `json:"nama_barang"`
	Tipe                   string  `json:"tipe"`
	TanggalPeroleh         string  `json:"tanggal_peroleh"`
	KondisiBarang          string  `json:"kondisi_barang"`
	SumberDana             string  `json:"sumber_dana"`
	KodeBarang             string  `json:"kode_barang"`
	HargaSatuan            float64 `json:"harga_satuan"`
	KodeInventaris         string  `json:"kode_inventaris"`
	RuanganID              int     `json:"ruangan_id"`
	Ruangan                Ruangan `json:"ruangan" gorm:"foreignKey:RuanganID"`
	NamaPengguna           string  `json:"nama_pengguna"`
	Unit                   string  `json:"unit"`
	Keterangan             string  `json:"keterangan"`
	Digunakan              int     `json:"digunakan"`
	TidakDigunakan         int     `json:"tidak_digunakan"`
	PengembalianLaptopLama string  `json:"pengembalian_laptop_lama"`
}
