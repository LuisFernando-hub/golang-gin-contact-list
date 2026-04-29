package utils

import (
	"errors"
	"fmt"
	"gin-contact-list/config"
	"gin-contact-list/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetContactById(c *gin.Context, contact *models.Contact) (int, int, error) {
	id, err := strconv.Atoi(c.Param("contactId"))

	if err != nil {
		return http.StatusBadRequest, id, errors.New("Invalid contact id")
	}

	if err := config.DB.Find(contact).Error; err != nil {
		return http.StatusInternalServerError, id, errors.New(fmt.Sprintf("Failed to find contact: %v", err.Error()))
	}

	if contact.ID == 0 {
		return http.StatusNotFound, id, errors.New(fmt.Sprintf("Contact with id %v not found", id))
	}

	return http.StatusOK, id, nil
}
