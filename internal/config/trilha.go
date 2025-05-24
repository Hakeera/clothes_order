package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfiguracaoTrilha map[string]map[string]PecaConfig

type PecaConfig struct {
	Modelos []string `json:"modelos"`
	Tipos   []string `json:"tipos"`
	Tecidos []string `json:"tecidos"`
}

var Trilha ConfiguracaoTrilha

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

