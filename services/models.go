// models.go
package services

import "time"

// Cliente representa a estrutura de dados de um cliente.
type Cliente struct {
	ID        int		`json:"c_id"`
	Nome      string	`json:"nome"`
	Email     string	`json:"email"`
	Telefone  string	`json:"telefone"`
	Historico string	`json:"historico"`
}

// Atendimento representa uma solicitação ou agendamento de atendimento.
type Atendimento struct {
	ID        int		`json:"a_id"`
	ClienteID int		`json:"c_id"`
	MedicoID  int		`json:"m_id"`
	Data      time.Time	`json:"data"`
	Descricao string	`json:"desc"`
	Status    string 	`json:"status"`// Ex: "Solicitado", "Agendado", "Realizado", "Cancelado"
}

// Documento representa documentos médicos, como exames e receitas.
type Documento struct {
	ID            int		`json:"d_id"`
	AtendimentoID int		`json:"a_id"`
	Tipo          string 	`json:"type"`// Ex: "Receita", "Resultado de Exame"
	Conteudo      string	`json:"content"`
	DataEmissao   time.Time	`json:"data_em"`
}

// --- Simulação de Banco de Dados ---
// Usaremos maps para armazenar os dados em memória.
var (
	clientes        = make(map[int]Cliente)
	atendimentos    = make(map[int]Atendimento)
	documentos      = make(map[int]Documento)
)