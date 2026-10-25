package handlers

import (
	"net/http"

	"my-website-backend/internal/models"
	"my-website-backend/internal/session"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) HandleLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	sessionID := c.Get("sessionID").(string)
	store := c.Get("sessionStore").(*session.Store)

	if email == "" || password == "" {
		return c.Render(http.StatusBadRequest, "login.html", map[string]interface{}{
			"Error": "Email and password are required.",
			"CSRF":  c.Get("csrf"),
		})
	}

	var user models.User
	result := h.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return c.Render(http.StatusUnauthorized, "login.html", map[string]interface{}{
			"Error": "Invalid email or password.",
			"CSRF":  c.Get("csrf"),
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return c.Render(http.StatusBadRequest, "login.html", map[string]interface{}{
			"Error": "Invalid email or password.",
			"CSRF":  c.Get("csrf"),
		})
	}

	store.Set(sessionID, "authenticated", true)
	store.Set(sessionID, "user_id", user.ID)
	store.Set(sessionID, "is_admin", user.IsAdmin)

	return c.Redirect(http.StatusSeeOther, "/admin/panel")
}

func (h *Handler) HandleLogout(c echo.Context) error {
	sessionID := c.Get("sessionID").(string)
	store := c.Get("sessionStore").(*session.Store)
	store.Delete(sessionID)

	return c.Redirect(http.StatusSeeOther, "/admin/login")
}
