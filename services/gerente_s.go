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
	fmt.Printf("Gerente cadastrou o(a) médico(a) '%s' (ID: %d).\n", m.Nome, m.ID)
	return nil
}

// LerMedico busca um médico pelo seu ID.
func LerMedico(id int) (Medico, bool) {
	m, ok := medicos[id]
	return m, ok
}

// AtualizarMedico atualiza as informações de um médico existente.
func AtualizarMedico(m Medico) error {
	if _, existe := medicos[m.ID]; !existe {
		return fmt.Errorf("médico com ID %d não encontrado para atualizar", m.ID)
	}

	medicos[m.ID] = m
	fmt.Printf("Gerente atualizou os dados do(a) médico(a) ID %d.\n", m.ID)
	return nil
}

// DeletarMedico remove um médico do sistema.
func DeletarMedico(id int) error {
	if _, existe := medicos[id]; !existe {
		return fmt.Errorf("médico com ID %d não encontrado para deletar", id)
	}

	delete(medicos, id)
	fmt.Printf("Gerente removeu o(a) médico(a) ID %d do sistema.\n", id)
	return nil
}

// ListarTodosMedicos retorna uma lista com todos os médicos cadastrados.
func ListarMedicos() []Medico {
	listaMedicos := make([]Medico, 0, len(medicos))
	for _, medico := range medicos {
		listaMedicos = append(listaMedicos, medico)
	}
	return listaMedicos
}
