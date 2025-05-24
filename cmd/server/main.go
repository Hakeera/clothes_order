package main

import (
	"clothes_order/internal/config"
	"clothes_order/internal/handlers"
	"clothes_order/internal/routes"
	"log"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func main() {
	// Carregar Json com Trilhas
	if err := config.CarregarTrilha("internal/config/trilha.json"); err != nil {
		log.Fatalf("erro ao carregar trilha: %v", err)
	}

	//Iniciar Server
	e := echo.New()

	// Configurar o renderizador de templates
	templateDir := filepath.Join("web", "templates")
	e.Renderer = handlers.NewTemplateRenderer(templateDir)
	e.Static("/static", "web/static")

	// Configurar Rotas
	routes.RegisterRoutes(e, e.Renderer)

	e.Logger.Fatal(e.Start(":1323"))
}

