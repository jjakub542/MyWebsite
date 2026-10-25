package app

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"

	_ "github.com/joho/godotenv/autoload"
)

type templateRegistry struct {
	templates map[string]*template.Template
}

func (t *templateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)

		return err
	}

	return tmpl.ExecuteTemplate(w, "base", data)
}

func Renderer() *templateRegistry {
	templates := make(map[string]*template.Template)

	templates["public/home.html"] = template.Must(template.ParseFiles("web/templates/public/home.html", "web/templates/public/base.html"))

	templates["admin/login.html"] = template.Must(template.ParseFiles("web/templates/admin/login.html", "web/templates/admin/base.html"))
	templates["admin/panel.html"] = template.Must(template.ParseFiles("web/templates/admin/panel.html", "web/templates/admin/base.html"))

	return &templateRegistry{
		templates: templates,
	}
}
