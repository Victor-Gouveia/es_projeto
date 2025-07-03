// atendimento.go
package services

//import "fmt"

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
func AtualizarAtendimento(id int, dadosAtualizados Atendimento) bool {
	if _, ok := atendimentos[id]; ok {
		// usa new_user para garantir que so atualizara espacos em branco
		// nao deve alterar o id
		var new_atend = atendimentos[id]
		if dadosAtualizados.ClienteID != 0 {
			new_atend.ClienteID = dadosAtualizados.ClienteID
		}
		if dadosAtualizados.MedicoID != 0 {
			new_atend.MedicoID = dadosAtualizados.MedicoID
		}
		if dadosAtualizados.Status != "" {
			new_atend.Status = dadosAtualizados.Status
		}
		if dadosAtualizados.Data != "" {
			new_atend.Data = dadosAtualizados.Data
		}
		if dadosAtualizados.Descricao != "" {
			new_atend.Descricao = dadosAtualizados.Descricao
		}
		atendimentos[id] = new_atend
		//fmt.Printf("Dados do atendimento %d atualizados.\n", id)
		return true
	}
	return false
}

// CancelarAtendimento (Delete) remove um atendimento.
func DeletarAtendimento(id int) bool {
	if _, ok := atendimentos[id]; ok {
		delete(atendimentos, id)
		//fmt.Printf("Atendimento %d removido.\n", id)
		return true
	}
	return false
}


