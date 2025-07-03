// cliente_ops.go
package services

import (
	"errors"
	"fmt"
)

// --- Use Case: Gerenciar Conta de Cliente (CRUD) ---

// CriarCliente adiciona um cliente com um ID definido manualmente.
func CriarCliente(c Cliente) error { // implementado
	// Validação: Garante que um ID foi fornecido e não está em uso.
	if c.ID == 0 {
		return errors.New("ID do cliente não pode ser 0")
	}
	if _, existe := clientes[c.ID]; existe {
		return fmt.Errorf("cliente com ID %d já existe", c.ID)
	}

	clientes[c.ID] = c
	fmt.Printf("Cliente '%s' (ID: %d) criado com sucesso.\n", c.Nome, c.ID)
	return nil
}

// LerCliente, AtualizarCliente, DeletarCliente, etc. permanecem os mesmos...
func LerCliente(id int) (Cliente, bool) { // implementado
	c, ok := clientes[id]
	return c, ok
}

// ListarClientes retorna uma lista de todos os clientes cadastrados.
func ListarClientes() []Cliente { // implementado
	// Cria um slice para armazenar os clientes.
	// O segundo argumento '0' é o tamanho inicial e o terceiro 'len(clientes)' é a capacidade inicial,
	// o que é uma pequena otimização.
	listaClientes := make([]Cliente, 0, len(clientes))

	// Itera sobre os valores do mapa 'clientes'. O '_' ignora a chave (o ID).
	for _, cliente := range clientes {
		listaClientes = append(listaClientes, cliente)
	}

	return listaClientes
}

func AtualizarCliente(id int, dadosAtualizados Cliente) bool {
	if _, ok := clientes[id]; ok {
		dadosAtualizados.ID = id
		clientes[id] = dadosAtualizados
		fmt.Printf("Dados do cliente %d atualizados.\n", id)
		EnviarDadosDoCliente(dadosAtualizados)
		return true
	}
	return false
}

func DeletarCliente(id int) bool {
	if _, ok := clientes[id]; ok {
		delete(clientes, id)
		fmt.Printf("Cliente %d deletado com sucesso.\n", id)
		return true
	}
	return false
}

// --- Use Case: Enviar Dados do Cliente (<<extends>> Gerenciar Conta) ---
func EnviarDadosDoCliente(c Cliente) {
	fmt.Printf(">> [Extensão] Enviando dados atualizados do cliente '%s' para sistemas externos...\n", c.Nome)
}

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
	// Define o status padrão para uma solicitação
	a.Status = "Solicitado"
	atendimentos[a.ID] = a
	fmt.Printf("Atendimento (ID: %d) solicitado pelo cliente %d.\n", a.ID, a.ClienteID)
	return nil
}

// --- Use Case: Visualizar Documentos ---
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
	fmt.Printf("Buscando documentos para o cliente %d...\n", clienteID)
	return docsCliente
}
