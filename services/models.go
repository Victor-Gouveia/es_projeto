// models.go
package services

// Cliente representa a estrutura de dados de um cliente.
type Cliente struct {
	ID        int    `json:"c_id"`
	Nome      string `json:"nome"`
	Email     string `json:"email"`
	Telefone  string `json:"telefone"`
	Historico string `json:"historico"`
}

// Atendimento representa uma solicitação ou agendamento de atendimento.
type Atendimento struct {
	ID        int    `json:"a_id"`
	ClienteID int    `json:"c_id"`
	MedicoID  int    `json:"m_id"`
	Data      string `json:"data"`
	Descricao string `json:"desc"`
	Status    string `json:"status"` // Ex: "Solicitado", "Agendado", "Realizado", "Cancelado"
}

// Documento representa documentos médicos, como exames e receitas.
type Documento struct {
	ID            int    `json:"d_id"`
	AtendimentoID int    `json:"a_id"`
	Tipo          string `json:"type"` // Ex: "Receita", "Resultado de Exame"
	Conteudo      string `json:"content"`
	DataEmissao   string `json:"data_em"`
}

type Medico struct {
	ID            int    `json:"m_id"`
	Nome          string `json:"nome"`
	CRM           string `json:"crm"`
	Especialidade string `json:"especialidade"`
	Telefone      string `json:"telefone"`
}

// --- Simulação de Banco de Dados ---
// Usaremos maps para armazenar os dados em memória.
var (
	// Mapa de clientes inicializado com 3 clientes.
	clientes = map[int]Cliente{
		1: {ID: 1, Nome: "João da Silva", Email: "joao.silva@email.com", Telefone: "11999991111"},
		2: {ID: 2, Nome: "Maria Oliveira", Email: "maria.o@email.com", Telefone: "21988882222"},
		3: {ID: 3, Nome: "Pedro Martins", Email: "pedro.m@email.com", Telefone: "84977773333"},
	}

	// Mapa de atendimentos inicializado com 3 atendimentos ligados aos clientes acima.
	atendimentos = map[int]Atendimento{
		101: {ID: 101, ClienteID: 1, MedicoID: 91, Data: "2025-07-27", Descricao: "Consulta de Rotina", Status: "Agendado"},
		102: {ID: 102, ClienteID: 2, MedicoID: 92, Data: "2025-07-04", Descricao: "Exame de Sangue", Status: "Realizado"},
		103: {ID: 103, ClienteID: 1, MedicoID: 93, Data: "2025-08-08", Descricao: "Exame de Sangue", Status: "Agendado"},
	}

	// Mapa de documentos inicializado com 1 documento ligado a um atendimento.
	documentos = map[int]Documento{
		5001: {ID: 5001, AtendimentoID: 102, Tipo: "Resultado de Exame", Conteudo: "Hemograma completo dentro dos padrões.", DataEmissao: "2025-07-19"},
		5002: {ID: 5002, AtendimentoID: 101, Tipo: "Receita", Conteudo: "Receita de Paracetamol 200mg", DataEmissao: "2025-07-27"},
	}

	medicos = map[int]Medico{
		91: {ID: 91, Nome: "Dr. Carlos Andrade", CRM: "RN 12345", Especialidade: "Cardiologia", Telefone: "84911112222"},
		92: {ID: 92, Nome: "Dra. Lucia Ferreira", CRM: "RN 54321", Especialidade: "Ortopedia", Telefone: "84933334444"},
		93: {ID: 93, Nome: "Dra. Marcia Lima", CRM: "RN 54321", Especialidade: "Clinico Geral", Telefone: "84933334444"},
	}
)
