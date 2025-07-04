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
	//fmt.Printf("Cliente '%s' (ID: %d) criado com sucesso.\n", c.Nome, c.ID)
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
		// usa new_user para garantir que so atualizara espacos em branco
		// nao deve alterar o id do usuario
		var new_user = clientes[id]
		if dadosAtualizados.Nome != "" {
			new_user.Nome = dadosAtualizados.Nome
		}
		if dadosAtualizados.Email != "" {
			new_user.Email = dadosAtualizados.Email
		}
		if dadosAtualizados.Telefone != "" {
			new_user.Telefone = dadosAtualizados.Telefone
		}
		if dadosAtualizados.Historico != "" {
			new_user.Historico = dadosAtualizados.Historico
		}
		clientes[id] = new_user
		//fmt.Printf("Dados do cliente %d atualizados.\n", id)
		return true
	}
	return false
}

func DeletarCliente(id int) bool {
	if _, ok := clientes[id]; ok {
		delete(clientes, id)
		//fmt.Printf("Cliente %d deletado com sucesso.\n", id)
		return true
	}
	return false
}
