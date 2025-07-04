// atendimento.go
package services

import "fmt"
import "errors"

// --- Use Case: Solicitar Atendimento ---
// O cliente agora "solicita" fornecendo os dados do atendimento, incluindo o ID desejado.
func SolicitarAtendimento(a Atendimento) error { // implementado
	if a.ID == 0 {
		return errors.New("ID do atendimento não pode ser 0")
	}
	if _, existe := atendimentos[a.ID]; existe {
		return fmt.Errorf("atendimento com ID %d já existe", a.ID)
	}
	if _, ok := LerCliente(a.ClienteID); !ok {
		return fmt.Errorf("cliente com ID %d nao encontrado", a.ClienteID)
	}
	if _, ok := LerMedico(a.MedicoID); !ok {
		return fmt.Errorf("medico com ID %d nao encontrado", a.MedicoID)
	}
	// Define o status padrão para uma solicitação
	a.Status = "Solicitado"
	atendimentos[a.ID] = a
	//fmt.Printf("Atendimento (ID: %d) solicitado pelo cliente %d.\n", a.ID, a.ClienteID)
	return nil
}

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


