// internal/jobs/handlers.go
package jobs

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var JobQueue chan Job

func Processar(c echo.Context) error  {
	tipo := c.Param("tipo")
	dado := c.FormValue("dado")

	job := Job{
		ID:   uuid.New().String(),
		Tipo: tipo,
		Dado: dado,
	}
	// Envia o job para a fila
	JobQueue <- job

	return c.JSON(http.StatusAccepted, map[string]string{
		"message": "Job enviado com sucesso",
		"id":      job.ID,
	})
}

func ProcessarCliente(job Job) {
	fmt.Printf("[Handler Cliente] Processando: %s\n", job.Dado)
}

func ProcessarProduto(job Job) {
	fmt.Printf("[Handler Produto] Processando: %s\n", job.Dado)
}

func ProcessarPedido(job Job) {
	fmt.Printf("[Handler Pedido] Processando: %s\n", job.Dado)
}

