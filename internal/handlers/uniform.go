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

func GetLinhas(c echo.Context) error {
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
