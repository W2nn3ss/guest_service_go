package repository

import "guest-service/models"

type GuestRepository interface {
	CreateGuest(guest *models.Guest) error
	GetGuestById(id uint) (models.Guest, error)
	GetAllGuests() ([]models.Guest, error)
	UpdateGuest(guest *models.Guest) error
	DeleteGuest(id uint) error
}
