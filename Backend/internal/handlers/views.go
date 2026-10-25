package handlers

import (
	"my-website-backend/internal/session"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ServeLoginPage(c echo.Context) error {
	sessionID := c.Get("sessionID").(string)
	store := c.Get("sessionStore").(*session.Store)

	if auth, ok := store.Get(sessionID, "authenticated"); ok && auth.(bool) {
		return c.Redirect(http.StatusSeeOther, "/admin/panel")
	}

	return c.Render(http.StatusOK, "admin/login.html", map[string]interface{}{
		"CSRF": c.Get("csrf"),
	})
}

func (h *Handler) ServeAdminPanel(c echo.Context) error {
	return c.Render(http.StatusOK, "admin/panel.html", map[string]interface{}{
		"CSRF": c.Get("csrf"),
	})
}

func (h *Handler) ServeHomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "public/home.html", map[string]interface{}{
		"CSRF": c.Get("csrf"),
	})
}

func (h *Handler) ServeSPA(c echo.Context) error {
	return c.File("web/static/dist/index.html")
}
