// atendente_ops.go
package services

import "fmt"

// --- Use Case: Gerenciar Atendimentos (CRUD) ---
// (Usado por Atendente e Gerente)

// AgendarAtendimento esta em cliente_s.go

// LerAtendimento (Read) busca um atendimento pelo ID.
func LerAtendimento(id int) (Atendimento, bool) {
	a, ok := atendimentos[id]
	return a, ok
}

// --- Use Case: Visualizar Atendimentos ---
// (Usado por Atendente e Gerente)
func ListarAtendimentos() []Atendimento { // implementado
	var lista []Atendimento
	for _, a := range atendimentos {
		lista = append(lista, a)
	}
	return lista
}

func LerAtendimentosCliente(c_id int) []Atendimento { // implementado
	// Cria uma lista para armazenar os atendimentos encontrados.
	var atendCliente []Atendimento

	// Itera sobre todos os atendimentos no mapa.
	for _, atendimento := range atendimentos {
		// Se o ID do cliente no atendimento for igual ao ID procurado...
		if atendimento.ClienteID == c_id {
			// ...adiciona o atendimento à lista.
			atendCliente = append(atendCliente, atendimento)
		}
	}

	// Retorna a lista de atendimentos encontrados (pode estar vazia se nenhum for encontrado).
	return atendCliente
}

func LerAtendimentosMedico(m_id int) []Atendimento { // implementado
	// Cria uma lista para armazenar os atendimentos encontrados.
	var atendCliente []Atendimento

	// Itera sobre todos os atendimentos no mapa.
	for _, atendimento := range atendimentos {
		// Se o ID do cliente no atendimento for igual ao ID procurado...
		if atendimento.MedicoID == m_id {
			// ...adiciona o atendimento à lista.
			atendCliente = append(atendCliente, atendimento)
		}
	}

	// Retorna a lista de atendimentos encontrados (pode estar vazia se nenhum for encontrado).
	return atendCliente
}

// AtualizarAtendimento (Update) atualiza um atendimento existente.
func AtualizarAtendimento(id int, status string, medicoID int) bool {
	if a, ok := atendimentos[id]; ok {
		a.Status = status
		a.MedicoID = medicoID
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


