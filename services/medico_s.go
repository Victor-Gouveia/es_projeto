// medico_ops.go
package services

import (
	"errors"
	"fmt"
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

func ListarDocumentos() []Documento { // implementado
	// Cria um slice (lista) para armazenar os documentos.
	// A capacidade inicial é definida como o tamanho do mapa para otimizar a performance.
	listaDocumentos := make([]Documento, 0, len(documentos))

	// Itera sobre os valores do mapa 'documentos'.
	// O '_' é usado para ignorar a chave (ID do documento), pois só queremos o objeto.
	for _, documento := range documentos {
		// Adiciona cada documento encontrado à lista.
		listaDocumentos = append(listaDocumentos, documento)
	}

	// Retorna a lista completa.
	return listaDocumentos
}

func ListarDocumentosMedico(medicoID int) []Documento { // implementado
	var docsMed []Documento
	for _, atendimento := range atendimentos {
		if atendimento.MedicoID == medicoID {
			for _, doc := range documentos {
				if doc.AtendimentoID == atendimento.ID {
					docsMed = append(docsMed, doc)
				}
			}
		}
	}
	fmt.Printf("Buscando documentos para o medico %d...\n", medicoID)
	return docsMed
}
