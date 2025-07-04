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
	if _, ok := LerAtendimento(d.AtendimentoID); !ok {
		return fmt.Errorf("atendimento com ID %d não existe", d.AtendimentoID)
	}

	documentos[d.ID] = d
	//fmt.Printf("Médico criou documento '%s' (ID: %d).\n", d.Tipo, d.ID)
	return nil
}

func LerDocumento(id int) (Documento, bool) {
	d, ok := documentos[id]
	return d, ok
}

// --- Use Case: Visualizar Documentos ---
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

func ListarDocumentosCliente(clienteID int) []Documento { // implementado
	var docsCliente []Documento
	for _, atendimento := range atendimentos {
		if atendimento.ClienteID == clienteID {
			for _, doc := range documentos {
				if doc.AtendimentoID == atendimento.ID {
					docsCliente = append(docsCliente, doc)
				}
			}
		}
	}
	//fmt.Printf("Buscando documentos para o cliente %d...\n", clienteID)
	return docsCliente
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
	//fmt.Printf("Buscando documentos para o medico %d...\n", medicoID)
	return docsMed
}

func AtualizarDocumento(id int, dadosAtualizados Documento) bool {
	if _, ok := documentos[id]; ok {
		// usa new_* para garantir que so atualizara espacos preenchidos
		// se nao fizer assim ele deixa espacos em branco
		// nao deve alterar o id
		var new_doc = documentos[id]
		if dadosAtualizados.AtendimentoID != 0 {
			new_doc.AtendimentoID = dadosAtualizados.AtendimentoID
		}
		if dadosAtualizados.Conteudo != "" {
			new_doc.Conteudo = dadosAtualizados.Conteudo
		}
		if dadosAtualizados.DataEmissao != "" {
			new_doc.DataEmissao = dadosAtualizados.DataEmissao
		}
		if dadosAtualizados.Tipo != "" {
			new_doc.Tipo = dadosAtualizados.Tipo
		}

		documentos[id] = new_doc
		//fmt.Printf("Dados do atendimento %d atualizados.\n", id)
		return true
	}
	return false
}

func DeletarDocumento(id int) bool {
	if _, ok := documentos[id]; ok {
		delete(documentos, id)
		//fmt.Printf("Médico deletou o documento %d.\n", id)
		return true
	}
	return false
}
