package controllers

import (
	"net/http"

	"phonebook/models"
	"phonebook/services"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type ContactController struct {
	Service *services.ContactService
}

func NewContactController(db *gorm.DB) *ContactController {
	service := services.NewContactService(db)
	return &ContactController{Service: service}
}

func (c *ContactController) CreateContact(ctx *gin.Context) {
	var contact models.Contact
	if err := ctx.ShouldBindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	if err := c.Service.CreateContact(&contact); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, contact)
}

func (c *ContactController) GetContacts(ctx *gin.Context) {
	pageStr := ctx.Query("page")
	limitStr := ctx.Query("limit")

	page := 1
	limit := 10

	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err == nil {
			page = p
		}
	}

	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err == nil && l <= 10 {
			limit = l
		}
	}

	contacts, err := c.Service.GetContacts(page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch contacts"})
		return
	}
	ctx.JSON(http.StatusOK, contacts)
}

func (c *ContactController) SearchContacts(ctx *gin.Context) {
	firstName := ctx.Query("first_name")
	lastName := ctx.Query("last_name")
	phoneNumber := ctx.Query("phone_number")
	address := ctx.Query("address")

	contacts, err := c.Service.SearchContacts(firstName, lastName, phoneNumber, address)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, contacts)
}

func (c *ContactController) UpdateContact(ctx *gin.Context) {
	idStr := ctx.Param("id")
	var contact models.Contact

	if err := ctx.ShouldBindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	updatedContact, err := c.Service.UpdateContact(uint(id), contact)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedContact)
}

func (c *ContactController) DeleteContact(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := c.Service.DeleteContact(uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
