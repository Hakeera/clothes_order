package handlers

import (
	"clothes_order/internal/config"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "base.html", nil)
}

func LinhasHandler(c echo.Context) error {
    linhas := make([]string, 0, len(config.Trilha))
    for linha := range config.Trilha {
        linhas = append(linhas, linha)
    }
    
    // Debug: verificar dados
    fmt.Printf("Linhas encontradas: %+v\n", linhas)
    
    // Verificar se o renderer existe
    if c.Echo().Renderer == nil {
        return c.String(http.StatusInternalServerError, "Renderer não configurado")
    }
    
    return c.Render(http.StatusOK, "components/linhas.html", echo.Map{
        "Linhas": linhas,
    })
}

func PecasHandler(c echo.Context) error {
	linhaSelecionada := c.QueryParam("linha")

	// Verifica se a linha existe na trilha carregada
	pecasMap, ok := config.Trilha[linhaSelecionada]
	if !ok {
		return c.String(http.StatusNotFound, "Linha não encontrada")
	}

	// Extrai os nomes das peças disponíveis na linha
	pecas := make([]string, 0, len(pecasMap))
	for nome := range pecasMap {
		pecas = append(pecas, nome)
	}

	// Renderiza o template com as peças da linha
	return c.Render(http.StatusOK, "components/pecas.html", echo.Map{
		"Pecas": pecas,
		"Linha": linhaSelecionada,
	})
}

func ModelosHandler(c echo.Context) error {
	linhaSelecionada := c.QueryParam("linha")
	pecaSelecionada := c.QueryParam("peca")

	// Verifica se a linha existe na trilha carregada
	pecasMap, ok := config.Trilha[linhaSelecionada]
	if !ok {
		return c.String(http.StatusNotFound, "Linha não encontrada")
	}

	// Verifica se a peça existe dentro da linha
	pecaConfig, ok := pecasMap[pecaSelecionada]
	if !ok {
		return c.String(http.StatusNotFound, "Peça não encontrada")
	}

	fmt.Println("MODELOS: ", pecaConfig.Modelos)
	return c.Render(http.StatusOK, "components/modelos.html", echo.Map{
		"Modelos": pecaConfig.Modelos,
		"Linha":   linhaSelecionada,
		"Peca":    pecaSelecionada,
	})
}

