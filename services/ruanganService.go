package services

import (
	"example/inventaris-api-v1/models"
	"example/inventaris-api-v1/repositories"
)

type RuangService struct {
	RuangRepo repositories.RuangRepository
}

func NewRuangService(ruangRepo *repositories.RuangRepository) *RuangService {
	return &RuangService{RuangRepo: *ruangRepo}
}

func (rs *RuangService) CreateRuang(ruang *models.Ruangan) error {
	return rs.RuangRepo.CreateRuang(ruang)
}

func (rs *RuangService) GetAllRuang() ([]models.Ruangan, error) {
	return rs.RuangRepo.GetAllRuang()
}

func (rs *RuangService) GetByID(id string) (*models.Ruangan, error) {
	return rs.RuangRepo.GetByID(id)
}

func (rs *RuangService) UpdateRuang(ruang *models.Ruangan, updatedAttributes map[string]interface{}) error {
	return rs.RuangRepo.UpdateRuang(ruang, updatedAttributes)
}

func (rs *RuangService) DeleteRuang(id string) error {
	return rs.RuangRepo.DeleteRuang(id)
}