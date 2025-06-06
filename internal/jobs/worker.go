// internal/jobs/worker.go
package jobs

import (
	"fmt"
	"time"
)

// StartDispatcher inicia N workers (goroutines) que processam jobs da fila
func StartDispatcher(jobQueue chan Job) {
	const numWorkers = 3
	for i := 1; i <= numWorkers; i++ {
		go StartWorker(i, jobQueue)
	}
}

func StartWorker(id int, queue <-chan Job) {
	for job := range queue {
		fmt.Printf("[Worker %d] Iniciando job: %s (%s)\n", id, job.ID, job.Tipo)
		time.Sleep(1 * time.Second) // simula demora

		switch job.Tipo {
		case "cliente":
			ProcessarCliente(job)
		case "produto":
			ProcessarProduto(job)
		case "pedido":
			ProcessarPedido(job)
		default:
			fmt.Printf("[Worker %d] Tipo desconhecido: %s\n", id, job.Tipo)
		}

		fmt.Printf("[Worker %d] Job %s finalizado.\n", id, job.ID)
	}
}

