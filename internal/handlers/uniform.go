package handlers

import (
	"clothes_order/internal/config"
	"fmt"
	"net/http"
	"slices"

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
	
	// Verifica se a linha existe
	if !config.VerificarLinhaExiste(linhaSelecionada) {
		return c.String(http.StatusNotFound, "Linha não encontrada")
	}
	
	// Obtém nomes das peças
	pecas, err := config.ObterNomesPecas(linhaSelecionada)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Erro ao obter peças: "+err.Error())
	}
	
	// Obtém todas as imagens
	imagens, err := config.ObterTodasImagensDaLinha(linhaSelecionada)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Erro ao obter imagens: "+err.Error())
	}
	
	// Obtém o map completo das peças para acessar imagens individuais
	pecasCompletas := config.Trilha[linhaSelecionada]
	
	fmt.Printf("Peças encontradas: %+v\n", pecas)
	fmt.Printf("Imagens encontradas: %+v\n", imagens)
	
	// Renderiza o template
	return c.Render(http.StatusOK, "components/pecas.html", echo.Map{
		"Pecas":          pecas,
		"Imagens":        imagens,
		"PecasCompletas": pecasCompletas,
		"Linha":          linhaSelecionada,
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

	return c.Render(http.StatusOK, "components/modelos.html", echo.Map{
		"Modelos": pecaConfig.Modelos,
		"Linha":   linhaSelecionada,
		"Peca":    pecaSelecionada,
	})
}
func TecidosHandler(c echo.Context) error {
	linhaSelecionada := c.QueryParam("linha")
	pecaSelecionada := c.QueryParam("peca")
	modeloSelecionado := c.QueryParam("modelo")

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

	// Verifica se o modelo existe no slice de modelos
	modeloValido := false
	// Looping para checar
	modeloValido = slices.Contains(pecaConfig.Modelos, modeloSelecionado)

	if !modeloValido {
		return c.String(http.StatusNotFound, "Modelo não encontrado")
	}

	// Renderiza o template com os tecidos disponíveis
	return c.Render(http.StatusOK, "components/tecidos.html", echo.Map{
		"Tecidos": pecaConfig.Tecidos,
		"Linha":   linhaSelecionada,
		"Peca":    pecaSelecionada,
		"Modelo":  modeloSelecionado,
	})
}

