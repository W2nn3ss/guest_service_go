package controllers

import (
	"github.com/gin-gonic/gin"
	"guest-service/models"
	"guest-service/service"
	"net/http"
	"strconv"
)

type GuestController struct {
	Service *service.GuestService
}

func (ctl *GuestController) CreateGuest(c *gin.Context) {
	var guest models.Guest
	if err := c.ShouldBindJSON(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctl.Service.CreateGuest(&guest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, guest)
}

func (ctl *GuestController) GetGuestById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Невалидный id гостя"})
		return
	}

	guest, err := ctl.Service.GetGuestById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Гость не найден"})
		return
	}

	c.JSON(http.StatusOK, guest)
}

func (ctl *GuestController) GetGuest(c *gin.Context) {
	var guests []models.Guest
	guests, err := ctl.Service.GetGuests()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Гости не найдены"})
	}

	c.JSON(http.StatusOK, guests)
}

func (ctl *GuestController) UpdateGuest(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Невалидный id гостя"})
		return
	}

	var guest models.Guest
	if err := c.ShouldBindJSON(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	guest.ID = uint(id)

	err = ctl.Service.UpdateGuest(&guest)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Гость не найден"})
		return
	}

	c.JSON(http.StatusOK, guest)
}

func (ctl *GuestController) DeleteGuest(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Невалидный id гостя"})
		return
	}

	err = ctl.Service.DeleteGuest(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Гость не найден"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Гость удалён"})
}
