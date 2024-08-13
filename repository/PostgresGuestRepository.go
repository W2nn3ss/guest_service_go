package repository

import (
	"gorm.io/gorm"
	"guest-service/models"
)

type PostgresGuestRepository struct {
	DB *gorm.DB
}

func (r *PostgresGuestRepository) CreateGuest(guest *models.Guest) error {
	return r.DB.Create(guest).Error
}

func (r *PostgresGuestRepository) GetGuestById(id uint) (models.Guest, error) {
	var guest models.Guest
	err := r.DB.First(&guest, id).Error
	return guest, err
}

func (r *PostgresGuestRepository) GetAllGuests() ([]models.Guest, error) {
	var guests []models.Guest
	err := r.DB.Find(&guests).Error
	return guests, err
}

func (r *PostgresGuestRepository) UpdateGuest(guest *models.Guest) error {
	return r.DB.Save(guest).Error
}

func (r *PostgresGuestRepository) DeleteGuest(id uint) error {
	return r.DB.Delete(&models.Guest{}, id).Error
}
