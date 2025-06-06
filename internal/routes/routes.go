package routes

import (
	"clothes_order/internal/handlers"
	"clothes_order/internal/jobs"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, renderer echo.Renderer) {
	e.Renderer = renderer

	// Página inicial
	e.GET("/", handlers.HomeHandler)
	e.GET("/linhas", handlers.LinhasHandler)
	e.GET("/pecas", handlers.PecasHandler)
	e.GET("/modelos", handlers.ModelosHandler)
	e.GET("/tecidos", handlers.TecidosHandler)
	e.GET("/grades", handlers.TecidosHandler)

	// Jobs
	e.POST("/processar/:tipo", jobs.Processar) 
}
