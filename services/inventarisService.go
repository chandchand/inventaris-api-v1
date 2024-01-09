package services

import (
	"example/inventaris-api-v1/models"
	"example/inventaris-api-v1/repositories"
)

type InventarisService struct {
	InventRepo repositories.InventRepository
}

func NewInventService(inventRepo *repositories.InventRepository) *InventarisService {
	return &InventarisService{InventRepo: *inventRepo}
}

func (i *InventarisService) CreateInventaris(data *models.Inventaris) error {
	return i.InventRepo.CreateInventaris(data)
}

func (i *InventarisService) GetAllInventaris(startDate, endDate, kondisiBarang, pengembalianLaptop string) ([]models.Inventaris, error) {
	return i.InventRepo.GetAllInventaris(startDate, endDate, kondisiBarang, pengembalianLaptop)
}
func (i *InventarisService) SearchData(keywords string) ([]models.Inventaris, error) {
	return i.InventRepo.SearchData(keywords)
}

func (i *InventarisService) GetByID(id string) (*models.Inventaris, error) {
	return i.InventRepo.GetByID(id)
}

func (i *InventarisService) UpdateInventaris(data *models.Inventaris, updatedAttributes map[string]interface{}) error {
	return i.InventRepo.UpdateInventaris(data, updatedAttributes)
}

func (i *InventarisService) DeleteInventaris(id string) error {
	return i.InventRepo.DeleteInventaris(id)
}