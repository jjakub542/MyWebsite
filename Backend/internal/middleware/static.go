package middleware

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

func StaticHTTP(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Request().URL.Path
		staticDir := "web/static/dist"

		// 1. Explicitly ignore backend-rendered routes
		// (Though Echo's routing order usually handles this, it's good practice)
		if path == "/login" || path == "/register" || path == "/admin" || path == "/user/home" || path == "/admin/panel" {
			return next(c)
		}

		// 2. Check if the exact file exists on disk
		filePath := filepath.Join(staticDir, path)
		if info, err := os.Stat(filePath); err == nil && !info.IsDir() {
			// File exists (e.g., /assets/rolldown-runtime.js). Pass to e.Static to serve it.
			return next(c)
		}

		// 3. PREVENT THE MIME TYPE ERROR:
		// If the request is for a specific file extension (like .js or .css) that is missing,
		// return a standard 404 Not Found. Do NOT return index.html for missing assets.
		if strings.HasSuffix(path, ".js") || strings.HasSuffix(path, ".css") || strings.HasSuffix(path, ".map") {
			return echo.ErrNotFound
		}

		// 4. If it's a navigation request (no file extension), return the SPA fallback
		return c.File(filepath.Join(staticDir, "index.html"))
	}
}
