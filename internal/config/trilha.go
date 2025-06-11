package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// ConfiguracaoTrilha representa toda a estrutura do JSON
type ConfiguracaoTrilha map[string]map[string]PecaConfig

// PecaConfig representa a configuração de cada peça individual
type PecaConfig struct {
	Img     []string `json:"img"`
	Modelos []string `json:"modelos"`
	Tecidos []string `json:"tecidos"`
	Grade   []string `json:"grade,omitempty"`
}

// PecaData é uma estrutura helper para passar dados organizados para o template
type PecaData struct {
	Nome     string
	Imagens  []string
	Modelos  []string
	Tecidos  []string
	Grade    []string
}

// Variável global para armazenar a trilha carregada
var Trilha ConfiguracaoTrilha

// CarregarTrilha carrega o arquivo JSON de configuração da trilha
func CarregarTrilha(caminho string) error {
	file, err := os.ReadFile(caminho)
	if err != nil {
		return fmt.Errorf("erro ao ler trilha: %w", err)
	}
	
	err = json.Unmarshal(file, &Trilha)
	if err != nil {
		return fmt.Errorf("erro ao parsear trilha: %w", err)
	}
	
	return nil
}

// ObterPecasDaLinha retorna as peças de uma linha específica organizadas
func ObterPecasDaLinha(linha string) ([]PecaData, error) {
	pecasMap, ok := Trilha[linha]
	if !ok {
		return nil, fmt.Errorf("linha '%s' não encontrada", linha)
	}
	
	var pecasData []PecaData
	
	for nomePeca, infoPeca := range pecasMap {
		peca := PecaData{
			Nome:    nomePeca,
			Imagens: infoPeca.Img,
			Modelos: infoPeca.Modelos,
			Tecidos: infoPeca.Tecidos,
			Grade:   infoPeca.Grade,
		}
		pecasData = append(pecasData, peca)
	}
	
	return pecasData, nil
}

// ObterNomesPecas retorna apenas os nomes das peças de uma linha
func ObterNomesPecas(linha string) ([]string, error) {
	pecasMap, ok := Trilha[linha]
	if !ok {
		return nil, fmt.Errorf("linha '%s' não encontrada", linha)
	}
	
	pecas := make([]string, 0, len(pecasMap))
	for nomePeca := range pecasMap {
		pecas = append(pecas, nomePeca)
	}
	
	return pecas, nil
}

// ObterTodasImagensDaLinha retorna todas as imagens de uma linha
func ObterTodasImagensDaLinha(linha string) ([]string, error) {
	pecasMap, ok := Trilha[linha]
	if !ok {
		return nil, fmt.Errorf("linha '%s' não encontrada", linha)
	}
	
	var todasImagens []string
	for _, infoPeca := range pecasMap {
		todasImagens = append(todasImagens, infoPeca.Img...)
	}
	
	return todasImagens, nil
}

// VerificarLinhaExiste verifica se uma linha existe na trilha
func VerificarLinhaExiste(linha string) bool {
	_, ok := Trilha[linha]
	return ok
}

// ObterLinhasDisponiveis retorna todas as linhas disponíveis
func ObterLinhasDisponiveis() []string {
	linhas := make([]string, 0, len(Trilha))
	for linha := range Trilha {
		linhas = append(linhas, linha)
	}
	return linhas
}
