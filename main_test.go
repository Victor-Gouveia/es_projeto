package main

import (
	//"net/http"

	"engsoft/auth"

	//"github.com/gin-gonic/gin"

	"engsoft/services"

	"strconv"
	"testing"
)

// ======================
// TESTES DE AUTENTICAÇÃO
// ======================

// Teste de login
func Test_Login(t *testing.T) {
	var user = "test"
	var pass = "test"
	msg, ok := auth.LogUser(user, pass)
	if msg == "" || !ok {
		t.Errorf(`Mensagem: %s \nOK: %s`, msg, strconv.FormatBool(ok)) // nao estava conseguindo mandar o bool direto, n sei pq
	}
}

// Teste de login com erro
func Test_LoginE(t *testing.T) {
	var user = "testaaa"
	var pass = "testaaa"
	msg, ok := auth.LogUser(user, pass)
	if msg == "" || ok {
		t.Errorf(`Program returned ok when there was an error`) // nao estava conseguindo mandar o bool direto, n sei pq
	}
}

// Teste de checagem de token
func Test_CheckToken(t *testing.T) {
	var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MCwicm9sZSI6ImdlcmVudCIsInVzZXJuYW1lIjoiZ2VyZW50In0.7ZW_vHF1S471nr2HUqqlLyk10uj-WuQkUL41vhqxhfY"
	msg, ok := auth.CheckToken(token, []string{"gerent"})
	if !ok || msg != "" {
		t.Errorf(`Mensagem: %s \nOK: %s`, msg, strconv.FormatBool(ok))
	}
}

// Teste de checagem de token passando um token invalido
func Test_CheckTokenE(t *testing.T) {
	var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MCwicm9sZSI6ImdlcmVudCIsInVzZXJuYW1lIjoiZ2VyZW50In0.7ZW_vHF1S471nr2HUqqlLyk"
	msg, ok := auth.CheckToken(token, []string{"gerent"})
	if ok || msg == "" {
		t.Errorf(`Did not send an error when it should`)
	}
}

// Teste de checagem de token passando um token sem permissao
func Test_CheckTokenE2(t *testing.T) {
	var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwicm9sZSI6InVzZXIiLCJ1c2VybmFtZSI6InRlc3QifQ.IYBdrfj_pLEjjKUYifp0wLLQZN3W-_Cj_uikcOgVmIY"
	msg, ok := auth.CheckToken(token, []string{"gerent"})
	if ok || msg != "unauthorized: token does not have required role" {
		t.Errorf(`Did not send an error when it should. msg: %s`, msg)
	}
}

// Teste de criacao de usuario
func Test_CreateUser(t *testing.T) {
	var name = "teste2"
	var pass = "teste2"
	var role = "user"
	var id = 5
	ok := auth.CreateUser(name, pass, role, id)
	if !ok {
		t.Errorf(`Failed to create user %s, pass=%s, role=%s, id=%d`, name, pass, role, id)
	}
}

// Teste de criacao de usuario com um cargo invalido
func Test_CreateUserE(t *testing.T) {
	var name = "teste2"
	var pass = "teste2"
	var role = "testando"
	var id = 5
	ok := auth.CreateUser(name, pass, role, id)
	if ok {
		t.Errorf(`Create a failed user %s, pass=%s, role=%s, id=%d`, name, pass, role, id)
	}
}

// Teste de geracao de token
func Test_GenToken(t *testing.T) {
	var user = "test"
	var role = "user"
	var id = 1
	token := auth.GenerateToken(user, role, id)
	if token == "" {
		t.Errorf(`Failed to generate token`)
	}
}

// ==================
// TESTES DE SERVICOS
// ==================

// --- cliente
// teste de criar cliente
func Test_CCriar(t *testing.T) {
	var newCli = services.Cliente{ID: 4, Nome: "José de Arimateia", Email: "jose.ari@email.com", Telefone: "11988885555"}
	if err := services.CriarCliente(newCli); err != nil {
		t.Errorf(`Failed to create client %s`, newCli.Nome)
	}
}

// teste de criar cliente com um id ja existente
func Test_CCriarE(t *testing.T) {
	var newCli = services.Cliente{ID: 2, Nome: "José de Arimateia", Email: "jose.ari@email.com", Telefone: "11988885555"}
	if err := services.CriarCliente(newCli); err == nil {
		t.Errorf(`Created a failed client %s at id %d`, newCli.Nome, newCli.ID)
	}
}

// teste de leitura de cliente
func Test_CLer(t *testing.T) {
	var id = 1
	_, ok := services.LerCliente(id)
	if !ok {
		t.Errorf(`Failed to read client %d`, id)
	}
}

// teste de leitura de cliente com um id inexistente
func Test_CLerE(t *testing.T) {
	var id = 10
	_, ok := services.LerCliente(id)
	if ok {
		t.Errorf(`Read an unexisting client %d`, id)
	}
}

// teste de leitura de cliente
func Test_CLista(t *testing.T) {
	lista := services.ListarClientes()
	if len(lista) == 0 { // map criado com 3 clientes, nao pode vir 0
		t.Errorf(`Failed to read clients from map`)
	}
}

// teste de atualizar cliente
func Test_CAtual(t *testing.T) {
	var id = 1
	var newCli = services.Cliente{ID: 1, Nome: "Geislaine Angelo", Email: "geis@email.com", Telefone: "11977775555"}
	if !services.AtualizarCliente(id, newCli) {
		t.Errorf(`Failed to update client %d`, id)
	}
}

// teste de atualizar cliente com um id inexistente
func Test_CAtualE(t *testing.T) {
	var id = 10
	var newCli = services.Cliente{ID: 1, Nome: "Geislaine Angelo", Email: "geis@email.com", Telefone: "11977775555"}
	if services.AtualizarCliente(id, newCli) {
		t.Errorf(`Update client that do not exist %d`, id)
	}
}

// teste de deletar cliente
func Test_CDel(t *testing.T) {
	var id = 3
	if !services.DeletarCliente(id) {
		t.Errorf(`Failed to delete client %d`, id)
	}
}

// teste de deletar cliente com id inexistente
func Test_CDelE(t *testing.T) {
	var id = 13
	if services.DeletarCliente(id) {
		t.Errorf(`Delete client that doesnt exist %d`, id)
	}
}

// atendimentos
// teste de solicitar atendimento
func Test_ACriar(t *testing.T) {
	var newAtend = services.Atendimento{ID: 104, ClienteID: 1, MedicoID: 91, Data: "2025-07-27", Descricao: "Consulta de Rotina", Status: "Solicitado"}
	if err := services.SolicitarAtendimento(newAtend); err != nil {
		t.Errorf(`Failed to create appointment %d`, newAtend.ID)
	}
}

// teste de solicitar atendimento com um cliente inexistente
func Test_ACriarE(t *testing.T) {
	var newAtend = services.Atendimento{ID: 104, ClienteID: 100, MedicoID: 91, Data: "2025-07-27", Descricao: "Consulta de Rotina", Status: "Solicitado"}
	if err := services.SolicitarAtendimento(newAtend); err == nil {
		t.Errorf(`Created an appointment with non existent client %d`, newAtend.ID)
	}
}

// teste de solicitar atendimento com um medico inexistente
func Test_ACriarE2(t *testing.T) {
	var newAtend = services.Atendimento{ID: 104, ClienteID: 1, MedicoID: 910, Data: "2025-07-27", Descricao: "Consulta de Rotina", Status: "Solicitado"}
	if err := services.SolicitarAtendimento(newAtend); err == nil {
		t.Errorf(`Created an appointment with non existent medic %d`, newAtend.ID)
	}
}

// teste de leitura de atendimento por cliente
func Test_ALerC(t *testing.T) {
	var c_id = 2
	var list = services.LerAtendimentosCliente(c_id)
	if len(list) == 0 {
		t.Errorf(`Failed to read appointments from client %d`, c_id)
	}
}

// teste de leitura de atendimento por cliente com um cliente inexistente
func Test_ALerCE(t *testing.T) {
	var c_id = 20
	var list = services.LerAtendimentosCliente(c_id)
	if len(list) > 0 {
		t.Errorf(`Read appointments from non existent client %d`, c_id)
	}
}

// teste de leitura de atendimento por medico
func Test_ALerM(t *testing.T) {
	var m_id = 91
	var lista = services.LerAtendimentosMedico(m_id)
	if len(lista) == 0 {
		t.Errorf(`Failed to read appointments from medic %d`, m_id)
	}
}

// teste de leitura de atendimento por medico com medico inexistente
func Test_ALerME(t *testing.T) {
	var m_id = 910
	var lista = services.LerAtendimentosMedico(m_id)
	if len(lista) > 0 {
		t.Errorf(`Read appointments from non existent medic %d`, m_id)
	}
}

// teste de leitura de atendimento
func Test_ALer(t *testing.T) {
	var id = 101
	_, ok := services.LerAtendimento(id)
	if !ok {
		t.Errorf(`Failed to read appointment %d`, id)
	}
}

// teste de leitura de atendimento com id inexistente
func Test_ALerE(t *testing.T) {
	var id = 1010
	_, ok := services.LerAtendimento(id)
	if ok {
		t.Errorf(`Read non existing appointment %d`, id)
	}
}

// teste de listagem de atendimentos
func Test_ALista(t *testing.T) {
	lista := services.ListarAtendimentos()
	if len(lista) == 0 { // map criado com 3 clientes, nao pode vir 0
		t.Errorf(`Failed to read appointments from map`)
	}
}

// teste de atualizar atendimentos
func Test_AAtual(t *testing.T) {
	var id = 102
	var newAtend = services.Atendimento{ID: 102, ClienteID: 1, MedicoID: 91, Data: "2025-07-27", Descricao: "Consulta de Rotina", Status: "Solicitado"}
	if !services.AtualizarAtendimento(id, newAtend) {
		t.Errorf(`Failed to update appointment %d`, id)
	}
}

// teste de atualizar atendimentos com um id inexistente
func Test_AAtualE(t *testing.T) {
	var id = 1020
	var newAtend = services.Atendimento{ID: 102, ClienteID: 1, MedicoID: 91, Data: "2025-07-27", Descricao: "Consulta de Rotina", Status: "Solicitado"}
	if services.AtualizarAtendimento(id, newAtend) {
		t.Errorf(`Update non existing appointment %d`, id)
	}
}

// teste de deletar atendimentos
func Test_ADel(t *testing.T) {
	var id = 101
	if !services.DeletarAtendimento(id) {
		t.Errorf(`Failed to delete appointment %d`, id)
	}
}

// teste de deletar atendimentos com um id inexistente
func Test_ADelE(t *testing.T) {
	var id = 1010
	if services.DeletarAtendimento(id) {
		t.Errorf(`Delete non existing appointment %d`, id)
	}
}

// documentos
// teste de criar documento
func Test_DCriar(t *testing.T) {
	var newDoc = services.Documento{ID: 5004, AtendimentoID: 102, Tipo: "Receita", Conteudo: "Receita de Aspirina 500mg", DataEmissao: "2025-07-31"}
	if err := services.CriarDocumento(newDoc); err != nil {
		t.Errorf(`Failed to create doccument %d`, newDoc.ID)
	}
}

// teste de criar documento com um id ja existente
func Test_DCriarE(t *testing.T) {
	var newDoc = services.Documento{ID: 5001, AtendimentoID: 101, Tipo: "Receita", Conteudo: "Receita de Aspirina 500mg", DataEmissao: "2025-07-31"}
	if err := services.CriarDocumento(newDoc); err == nil {
		t.Errorf(`Created invalid doccument %d`, newDoc.ID)
	}
}

// teste de criar documento com um atendimento invalido
func Test_DCriarE2(t *testing.T) {
	var newDoc = services.Documento{ID: 5001, AtendimentoID: 1010, Tipo: "Receita", Conteudo: "Receita de Aspirina 500mg", DataEmissao: "2025-07-31"}
	if err := services.CriarDocumento(newDoc); err == nil {
		t.Errorf(`Created invalid doccument %d`, newDoc.ID)
	}
}

// teste de leitura de documento por cliente
func Test_DLerC(t *testing.T) {
	var c_id = 1
	var list = services.ListarDocumentosCliente(c_id)
	if len(list) == 0 {
		t.Errorf(`Failed to read doccuments from client %d`, c_id)
	}
}

// teste de leitura de documento por cliente nao existente
func Test_DLerCE(t *testing.T) {
	var c_id = 10
	var list = services.ListarDocumentosCliente(c_id)
	if len(list) > 0 {
		t.Errorf(`Read invalid doccuments from client %d`, c_id)
	}
}

// teste de leitura de atendimento por medico
func Test_DLerM(t *testing.T) {
	var m_id = 910
	var lista = services.ListarDocumentosMedico(m_id)
	if len(lista) > 0 {
		t.Errorf(`Failed to read doccuments from medic %d`, m_id)
	}
}

// teste de leitura de atendimento por medico com um id invalido
func Test_DLerME(t *testing.T) {
	var m_id = 910
	var lista = services.ListarDocumentosMedico(m_id)
	if len(lista) > 0 {
		t.Errorf(`Read invalid doccuments from medic %d`, m_id)
	}
}

// teste de leitura de atendimento
func Test_DLer(t *testing.T) {
	var id = 5001
	_, ok := services.LerDocumento(id)
	if !ok {
		t.Errorf(`Failed to read doccument %d`, id)
	}
}

// teste de leitura de atendimento invalido
func Test_DLerE(t *testing.T) {
	var id = 50010
	_, ok := services.LerDocumento(id)
	if ok {
		t.Errorf(`Read invalid doccument %d`, id)
	}
}

// teste de listagem de atendimentos
func Test_DLista(t *testing.T) {
	lista := services.ListarDocumentos()
	if len(lista) == 0 { // map criado maior q 0
		t.Errorf(`Failed to read doccuments from map`)
	}
}

// teste de atualizar atendimentos
func Test_DAtual(t *testing.T) {
	var id = 5001
	var newDoc = services.Documento{ID: 5002, AtendimentoID: 101, Tipo: "Receita", Conteudo: "Receita de Aspirina 500mg", DataEmissao: "2025-07-31"}
	if !services.AtualizarDocumento(id, newDoc) {
		t.Errorf(`Failed to update doccument %d`, id)
	}
}

// teste de atualizar atendimentos com id invalido
func Test_DAtualE(t *testing.T) {
	var id = 50010
	var newDoc = services.Documento{ID: 5002, AtendimentoID: 101, Tipo: "Receita", Conteudo: "Receita de Aspirina 500mg", DataEmissao: "2025-07-31"}
	if services.AtualizarDocumento(id, newDoc) {
		t.Errorf(`Updated invalid doccument %d`, id)
	}
}

// teste de deletar atendimentos
func Test_DDel(t *testing.T) {
	var id = 5002
	if !services.DeletarDocumento(id) {
		t.Errorf(`Failed to delete doccument %d`, id)
	}
}

// teste de deletar atendimentos com um id invalido
func Test_DDelE(t *testing.T) {
	var id = 50020
	if services.DeletarDocumento(id) {
		t.Errorf(`Deleted invalid doccument %d`, id)
	}
}

// --- medico
// teste de criar medico
func Test_MCriar(t *testing.T) {
	var newMed = services.Medico{ID: 94, Nome: "Dra. Emilly Miller", CRM: "RN 15243", Especialidade: "Clinico Geral", Telefone: "84933334444"}
	if err := services.CriarMedico(newMed); err != nil {
		t.Errorf(`Failed to create medic %s`, newMed.Nome)
	}
}

// teste de criar medico com um id ocupado
func Test_MCriarE(t *testing.T) {
	var newMed = services.Medico{ID: 91, Nome: "Dra. Emilly Miller", CRM: "RN 15243", Especialidade: "Clinico Geral", Telefone: "84933334444"}
	if err := services.CriarMedico(newMed); err == nil {
		t.Errorf(`Create invalid medic %s`, newMed.Nome)
	}
}

// teste de leitura de medico
func Test_MLer(t *testing.T) {
	var id = 91
	_, ok := services.LerMedico(id)
	if !ok {
		t.Errorf(`Failed to read medic %d`, id)
	}
}

// teste de leitura de medico com um id invalido
func Test_MLerE(t *testing.T) {
	var id = 910
	_, ok := services.LerMedico(id)
	if ok {
		t.Errorf(`Read invalid medic %d`, id)
	}
}

// teste de leitura de medico
func Test_MLista(t *testing.T) {
	lista := services.ListarMedicos()
	if len(lista) == 0 { // map criado maior q 0
		t.Errorf(`Failed to read medics from map`)
	}
}

// teste de atualizar medico
func Test_MAtual(t *testing.T) {
	var id = 92
	var newMed = services.Medico{ID: 94, Nome: "Dra. Emilly Miller", CRM: "RN 15243", Especialidade: "Clinico Geral", Telefone: "84933334444"}
	if !services.AtualizarMedico(id, newMed) {
		t.Errorf(`Failed to update medic %d`, id)
	}
}

// teste de atualizar medico com id invalido
func Test_MAtualE(t *testing.T) {
	var id = 920
	var newMed = services.Medico{ID: 94, Nome: "Dra. Emilly Miller", CRM: "RN 15243", Especialidade: "Clinico Geral", Telefone: "84933334444"}
	if services.AtualizarMedico(id, newMed) {
		t.Errorf(`updated invalid medic %d`, id)
	}
}

// teste de deletar medico
func Test_MDel(t *testing.T) {
	var id = 93
	if !services.DeletarMedico(id) {
		t.Errorf(`Failed to delete medic %d`, id)
	}
}

// teste de deletar medico invalido
func Test_MDelE(t *testing.T) {
	var id = 93
	if services.DeletarMedico(id) {
		t.Errorf(`Deleted invalid medic %d`, id)
	}
}
