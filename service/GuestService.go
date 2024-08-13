package service

import (
	"github.com/nyaruka/phonenumbers"
	"guest-service/models"
	"guest-service/repository"
	"log"
)

type GuestService struct {
	Repo repository.GuestRepository
}

func getCountryFromPhoneNumber(phone string) (string, error) {
	num, err := phonenumbers.Parse(phone, "")
	if err != nil {
		return "", err
	}
	region := phonenumbers.GetRegionCodeForNumber(num)
	return region, nil
}

func (s *GuestService) CreateGuest(guest *models.Guest) error {
	// Определяем страну по номеру телефона, если она не указана
	if guest.Country == "" {
		country, err := getCountryFromPhoneNumber(guest.Phone)
		if err != nil {
			log.Printf("Error parsing phone number: %v", err)
			return err
		}
		guest.Country = country
	}

	// Сохраняем гостя в базу данных
	if err := s.Repo.CreateGuest(guest); err != nil {
		return err
	}

	return nil
}

func (s *GuestService) GetGuestById(id uint) (models.Guest, error) {
	return s.Repo.GetGuestById(id)
}

func (s *GuestService) GetGuests() ([]models.Guest, error) {
	return s.Repo.GetAllGuests()
}

func (s *GuestService) UpdateGuest(guest *models.Guest) error {
	// Извлекаем существующего гостя из базы данных
	existingGuest, err := s.Repo.GetGuestById(guest.ID)
	if err != nil {
		return err
	}

	// Обновляем данные гостя
	existingGuest.FirstName = guest.FirstName
	existingGuest.LastName = guest.LastName
	existingGuest.Email = guest.Email
	existingGuest.Phone = guest.Phone

	if guest.Country == "" || existingGuest.Phone != guest.Phone {
		country, err := getCountryFromPhoneNumber(guest.Phone)
		if err != nil {
			log.Printf("Error parsing phone number: %v", err)
			return err
		}
		existingGuest.Country = country
	} else {
		existingGuest.Country = guest.Country
	}

	// Сохраняем обновленного гостя в базу данных
	if err := s.Repo.UpdateGuest(&existingGuest); err != nil {
		return err
	}

	return nil
}

func (s *GuestService) DeleteGuest(id uint) error {
	return s.Repo.DeleteGuest(id)
}
