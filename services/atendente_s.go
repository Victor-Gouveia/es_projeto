// atendente_ops.go
package services

import "fmt"
import "time"

// --- Use Case: Gerenciar Atendimentos (CRUD) ---
// (Usado por Atendente e Gerente)

// AgendarAtendimento esta em cliente_s.go

// LerAtendimento (Read) busca um atendimento pelo ID.
func LerAtendimento(id int) (Atendimento, bool) {
	a, ok := atendimentos[id]
	return a, ok
}

// AtualizarAtendimento (Update) atualiza um atendimento existente.
func AtualizarAtendimento(id int, status string, medicoID int) bool {
	if a, ok := atendimentos[id]; ok {
		a.Status = status
		a.MedicoID = medicoID
		if status == "Agendado" {
			a.Data = time.Now().Add(24 * time.Hour) // Reagenda para o dia seguinte
		}
		atendimentos[id] = a
		fmt.Printf("Atendimento %d atualizado para status '%s'.\n", id, status)
		return true
	}
	return false
}

// CancelarAtendimento (Delete) remove um atendimento.
func CancelarAtendimento(id int) bool {
	if _, ok := atendimentos[id]; ok {
		delete(atendimentos, id)
		fmt.Printf("Atendimento %d cancelado.\n", id)
		return true
	}
	return false
}

// --- Use Case: Visualizar Atendimentos ---
// (Usado por Atendente e Gerente)
func VisualizarTodosAtendimentos() []Atendimento {
	var lista []Atendimento
	for _, a := range atendimentos {
		lista = append(lista, a)
	}
	fmt.Println("Visualizando todos os atendimentos no sistema...")
	return lista
}