package repositories

import (
	"errors"
	"example/inventaris-api-v1/models"

	"gorm.io/gorm"
)

type RuangRepository struct {
	DB *gorm.DB
}

func NewRuangRepo(db *gorm.DB) *RuangRepository{
	return &RuangRepository{DB: db}
}

func (r *RuangRepository) CreateRuang(ruang *models.Ruangan) error {
	return r.DB.Create(ruang).Error
}

func (r *RuangRepository) GetAllRuang() ([]models.Ruangan, error) {
	var data []models.Ruangan
	if err := r.DB.Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("tidak ada data ruangan")
		}
	}
	return data, nil
}

func (r *RuangRepository) GetByID(id string) (*models.Ruangan, error) {
	var ruang models.Ruangan
	if err := r.DB.First(&ruang, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("tidak ada data ruangan dari id: " + id)
		}
		return nil, err
	}
	return &ruang, nil
}

func (r *RuangRepository) UpdateRuang(ruang *models.Ruangan, updatedAttributes map[string]interface{}) error {
	return r.DB.Model(ruang).Updates(updatedAttributes).Error
}

func (r *RuangRepository) DeleteRuang(id string) error {
	return r.DB.Delete(&models.Ruangan{}, "id = ?", id).Error
}