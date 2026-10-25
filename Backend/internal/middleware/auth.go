package middleware

import (
	"my-website-backend/internal/session"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionID := c.Get("sessionID").(string)
		store := c.Get("sessionStore").(*session.Store)
		auth, ok := store.Get(sessionID, "authenticated")
		if !ok || !auth.(bool) {
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		return next(c)
	}
}

func AdminAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionID := c.Get("sessionID").(string)
		store := c.Get("sessionStore").(*session.Store)
		isAdmin, ok := store.Get(sessionID, "is_admin")
		if !ok || !isAdmin.(bool) {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "Access forbidden: Admin only"})
		}
		return next(c)
	}
}
