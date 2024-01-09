package repositories

import (
	"errors"
	"example/inventaris-api-v1/models"

	"gorm.io/gorm"
)

type InventRepository struct {
	DB *gorm.DB
}

func NewInventRepo(db *gorm.DB) *InventRepository{
	return &InventRepository{DB: db}
}

func (ir *InventRepository) CreateInventaris(ruang *models.Inventaris) error {
	return ir.DB.Create(ruang).Error
}

func (ir *InventRepository) GetAllInventaris(startDate, endDate, kondisiBarang, pengembalianLaptop string) ([]models.Inventaris, error) {
	var data []models.Inventaris

	query := ir.DB.Find(&data).
		Preload("Ruangan").
		Where("1 = 1")


	if startDate != "" && endDate != ""{
		query = query.Where("tanggal_peroleh BETWEEN ? AND ?", startDate, endDate)
	}

	if kondisiBarang != "" {
		query = query.Where("kondisi_barang = ? ", kondisiBarang)
	}

	if pengembalianLaptop != "" {
		query = query.Where("pengembalian_laptop_lama = ? ", pengembalianLaptop)
	}

	if err := query.Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("tidak ada data inventaris")
		}
	}

	return data, nil
}

func (ir *InventRepository) SearchData(keywords string) ([]models.Inventaris, error) {
	var data []models.Inventaris

	query := ir.DB.Find(&data).
		Preload("Ruangan").
		Where("nama_barang LIKE ? OR kode_barang LIKE ? OR tipe LIKE ? OR kode_inventaris LIKE ? OR nama_pengguna LIKE ? OR unit LIKE ? OR kondisi_barang LIKE ? ", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%")

	if err := query.Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("tidak ada data pencarian nama barang dari " + keywords)
		}
	}

	return data, nil
}

func (ir *InventRepository) GetByID(id string) (*models.Inventaris, error) {
	var data models.Inventaris
	if err := ir.DB.Preload("Ruangan").First(&data, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("tidak ada data inventaris dari id: " + id)
		}
		return nil, err
	}
	return &data, nil
}

func (ir *InventRepository) UpdateInventaris(data *models.Inventaris, updatedAttributes map[string]interface{}) error {
	return ir.DB.Model(data).Updates(updatedAttributes).Error
}

func (ir *InventRepository) DeleteInventaris(id string) error {
	return ir.DB.Delete(&models.Inventaris{}, "id = ?", id).Error
}