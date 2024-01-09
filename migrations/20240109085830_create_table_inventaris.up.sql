CREATE TABLE IF NOT EXISTS inventaris (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama_barang VARCHAR(255),
    tipe VARCHAR(255),
    tanggal_peroleh DATE,
    kondisi_barang ENUM('Baik', 'Rusak Ringan', 'Rusak Berat'),
    sumber_dana VARCHAR(255),
    kode_barang VARCHAR(255),
    harga_satuan DECIMAL(10, 2),
    kode_inventaris VARCHAR(255),
    ruangan_id INT,
    nama_pengguna VARCHAR(255),
    unit VARCHAR(255),
    keterangan TEXT,
    digunakan INT,
    tidak_digunakan INT,
    pengembalian_laptop_lama ENUM('Sudah', 'Belum')
)ENGINE=InnoDB;