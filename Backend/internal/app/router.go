package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"

	"my-website-backend/internal/handlers"
	"my-website-backend/internal/middleware"
)

func (s *Server) Router() http.Handler {
	e := echo.New()
	e.Use(m.CSRFWithConfig(m.CSRFConfig{
		TokenLookup:  "header:X-CSRF-Token,form:_csrf",
		CookieName:   "_csrf",
		CookieSecure: false, // set to true if using HTTPS
	}))
	e.Use(middleware.Session(s.store))
	e.Renderer = Renderer()
	e.Static("/static", "web/static")
	e.Static("/assets", "web/static/dist/assets")
	e.Use(m.RequestLogger())
	e.Use(m.Recover())

	h := &handlers.Handler{DB: s.db}

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, "/home")
	})
	e.GET("/home", h.ServeHomePage)
	e.POST("/message", h.HandleSubmitMessage)

	e.GET("/admin/login", h.ServeLoginPage)
	e.POST("/admin/login", h.HandleLogin)
	e.POST("/admin/logout", h.HandleLogout)

	adminGroup := e.Group("/admin")
	adminGroup.Use(middleware.AdminAuth)
	adminGroup.GET("/panel", h.ServeAdminPanel)

	appGroup := e.Group("/app/engineering-thesis")
	appGroup.GET("", h.ServeSPA)
	appGroup.GET("/*", h.ServeSPA)

	return e
}
