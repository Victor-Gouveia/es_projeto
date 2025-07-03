package services

import (
	"errors"
	"fmt"
)

// CriarMedico adiciona um novo médico ao sistema.
func CriarMedico(m Medico) error {
	if m.ID == 0 {
		return errors.New("ID do médico não pode ser 0")
	}
	if _, existe := medicos[m.ID]; existe {
		return fmt.Errorf("médico com ID %d já existe", m.ID)
	}

	medicos[m.ID] = m
	//fmt.Printf("Gerente cadastrou o(a) médico(a) '%s' (ID: %d).\n", m.Nome, m.ID)
	return nil
}

// LerMedico busca um médico pelo seu ID.
func LerMedico(id int) (Medico, bool) {
	m, ok := medicos[id]
	return m, ok
}

// AtualizarMedico atualiza as informações de um médico existente.
func AtualizarMedico(id int, dadosAtualizados Medico) bool {
	if _, ok := medicos[id]; ok {
		// usa new_* para garantir que so atualizara espacos em branco
		// nao deve alterar o id do usuario
		var new_med = medicos[id]
		if dadosAtualizados.Nome != "" {
			new_med.Nome = dadosAtualizados.Nome
		}
		if dadosAtualizados.Telefone != "" {
			new_med.Telefone = dadosAtualizados.Telefone
		}
		if dadosAtualizados.CRM != "" {
			new_med.CRM = dadosAtualizados.CRM
		}
		if dadosAtualizados.Especialidade != "" {
			new_med.Especialidade = dadosAtualizados.Especialidade
		}
		medicos[id] = new_med
		//fmt.Printf("Dados do medico %d atualizados.\n", id)
		return true
	}
	return false
}

// DeletarMedico remove um médico do sistema.
func DeletarMedico(id int) bool {
	if _, ok := medicos[id]; ok {
		delete(medicos, id)
		//fmt.Printf("Medico %d deletado com sucesso.\n", id)
		return true
	}
	return false
}

// ListarTodosMedicos retorna uma lista com todos os médicos cadastrados.
func ListarMedicos() []Medico {
	listaMedicos := make([]Medico, 0, len(medicos))
	for _, medico := range medicos {
		listaMedicos = append(listaMedicos, medico)
	}
	return listaMedicos
}

// MaiorIDMedico percorre o mapa de médicos e retorna o maior ID encontrado.
func MaiorIDMedico() int {
	if len(medicos) == 0 {
		return 0
	}
	
	maiorID := 0
	for id := range medicos {
		if id > maiorID {
			maiorID = id
		}
	}
	return maiorID
}