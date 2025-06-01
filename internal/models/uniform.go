package models

type Uniform struct {
    ID           int    `json:"id"`
    Linha        string `json:"linha"`
    TipoPeca     string `json:"tipo_peca"`
    Modelo       string `json:"modelo"`
    Tecido       string `json:"tecido"`
    TipoDesenho  string `json:"tipo_desenho"`
    Posicao      string `json:"posicao"`
    LogoPath     string `json:"logo_path"`
    Preco        float64 `json:"preco"`
}
