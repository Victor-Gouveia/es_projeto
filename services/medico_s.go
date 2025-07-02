// medico_ops.go
package services

import (
	"errors"
	"fmt"
	"time"
)

// --- Use Case: Gerenciar Documentos (CRUD) ---

// CriarDocumento adiciona um novo documento com ID manual.
func CriarDocumento(d Documento) error {
	if d.ID == 0 {
		return errors.New("ID do documento não pode ser 0")
	}
	if _, existe := documentos[d.ID]; existe {
		return fmt.Errorf("documento com ID %d já existe", d.ID)
	}

	d.DataEmissao = time.Now()
	documentos[d.ID] = d
	fmt.Printf("Médico criou documento '%s' (ID: %d).\n", d.Tipo, d.ID)
	return nil
}

func LerDocumento(id int) (Documento, bool) {
	d, ok := documentos[id]
	return d, ok
}

func AtualizarDocumento(id int, conteudo string) bool {
	if d, ok := documentos[id]; ok {
		d.Conteudo = conteudo
		documentos[id] = d
		fmt.Printf("Médico atualizou o documento %d.\n", id)
		return true
	}
	return false
}

func DeletarDocumento(id int) bool {
	if _, ok := documentos[id]; ok {
		delete(documentos, id)
		fmt.Printf("Médico deletou o documento %d.\n", id)
		return true
	}
	return false
}

// --- Use Case: Visualizar Atendimentos ---
// O médico visualiza apenas os atendimentos atribuídos a ele.
func VisualizarAtendimentosMedico(medicoID int) []Atendimento {
	var lista []Atendimento
	for _, a := range atendimentos {
		if a.MedicoID == medicoID {
			lista = append(lista, a)
		}
	}
	fmt.Printf("Médico %d visualizando seus atendimentos...\n", medicoID)
	return lista
}
