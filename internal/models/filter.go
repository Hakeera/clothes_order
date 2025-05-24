package models

type FilterConfig struct {
    Linhas map[string][]string `json:"linhas"`
    Pecas  map[string][]string `json:"pecas"`
    Modelos map[string][]string `json:"modelos"`
    Tecidos []string `json:"tecidos"`
    Desenhos []string `json:"desenhos"`
    Posicoes []string `json:"posicoes"`
}
