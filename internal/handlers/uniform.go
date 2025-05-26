package handlers

import (
	"clothes_order/internal/config"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	// Pode passar dados adicionais no segundo argumento, se desejar
	return c.Render(http.StatusOK, "base.html", nil)
}

func LinhasHandler(c echo.Context) error {
    // Debug: verificar se está chegando aqui
    fmt.Println("GetLinhas chamada")
    
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
	fmt.Println("Linha:", linhaSelecionada)

	pecasMap := config.Trilha[linhaSelecionada]
	for peca := range pecasMap {
		
	// Verifica se a linha existe na trilha
	pecas, ok := config.Trilha[linhaSelecionada]
	if !ok {
		return c.String(http.StatusNotFound, "Linha não encontrada")
	}

	// Debug: verificar dados
	fmt.Printf("Peças encontradas: %+v\n", pecas)

	return c.Render(http.StatusOK, "components/pecas.html", echo.Map{
	"Peças": pecas ,
	})
}

