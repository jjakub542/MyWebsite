package handlers

import (
	"my-website-backend/internal/models"

	"gorm.io/gorm"

	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) HandleSubmitMessage(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	content := c.FormValue("content")

	if name == "" || email == "" || content == "" {
		return c.String(http.StatusBadRequest, "All fields are required")
	}

	msg := models.Message{
		Name:    name,
		Email:   email,
		Content: content,
	}

	if err := h.DB.Create(&msg).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save message")
	}

	return c.String(http.StatusOK, "Message sent successfully!")
}
