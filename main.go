package main

import (
	"net/http"

	"engsoft/auth"

	"github.com/gin-gonic/gin"

	"engsoft/services"

	"strconv"
)

// struct com informacoes dos albums, tudo com parametros json

func main() {
	router := gin.Default()

	//Rotas de autenticacao
	router.POST("/auth", postAuth)
	router.POST("/auth/new", postAuthNew)

	//Funcoes da aplicacao final real oficial
	//GET
	router.GET("/clientes", getClientes)
	router.GET("/clientes/:id", getClientesID)

	router.GET("/medicos", getMedicos)
	router.GET("/medicos/:id", getMedicosID)

	router.GET("/atendimentos", getAtends)
	router.GET("/atendimentos/cliente/:id", getAtendsCliente)
	router.GET("/atendimentos/medico/:id", getAtendsMedico)

	router.GET("/documentos", getDocs)
	router.GET("/documentos/cliente/:id", getDocsCliente)
	router.GET("/documentos/medico/:id", getDocsMedico)
	//POST
	router.POST("/clientes", postCliente)
	router.POST("/atendimentos", postAtend)
	router.POST("/documentos", postDoc)
	router.POST("/medicos", postMedico)
	//PUT
	router.PUT("/clientes/:id", putCliente)
	router.PUT("/medicos/:id", putMedico)
	router.PUT("/atendimentos/:id", putAtend)
	router.PUT("/documentos/:id", putDoc)
	//DELETE
	router.DELETE("/clientes/:id", delCliente)
	router.DELETE("/medicos/:id", delMedico)
	router.DELETE("/atendimentos/:id", delAtend)
	router.DELETE("/documentos/:id", delDoc)

	// roda o servidor ouvindo todos os ips
	router.Run("0.0.0.0:8080")

	// OBS: Existe um meio de testar o token ao receber uma rota,
	// mas por agora (12/06) nao irei fazer isso
	// VVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVV
	// admin := router.group("/")
	// admin.Use(*Funcao para autenticar*){*rotas de admin aqui*}
}

// Funcao para listar clientes
func getClientes(c *gin.Context) {
	if ok, msg := checkToken(c, []string{}); !ok{ // qualquer um pode pesquisar clientes (passivel de alteracao)
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}
	var clientes = services.ListarClientes()
	c.IndentedJSON(http.StatusOK, clientes)
}

// lista cliente pelo id
func getClientesID(c *gin.Context) {

	if ok, msg := checkToken(c, []string{}); !ok{ // qualquer um pode pesquisar cliente por id
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}
	id := c.Param("id")
	c_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "client id not valid"})
		return
	}
	cliente, ok := services.LerCliente(c_id)
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "client not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, cliente)
}

func postCliente(c *gin.Context) {
	var newCliente services.Cliente

	if ok, msg := checkToken(c, []string{"user", "atende", "gerent"}); !ok{ // medicos nao podem enviar dados do cliente
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}

	if err := c.BindJSON(&newCliente); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to obtain client"})
		return
	}

	if err := services.CriarCliente(newCliente); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newCliente)
}

func putCliente(c *gin.Context) {
	if ok, msg := checkToken(c, []string{"user", "atende", "gerent"}); !ok{ // medicos nao podem atualizar dados do cliente
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}

	id := c.Param("id")
	c_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "client id not valid"})
		return
	}
	var newCliente services.Cliente
	if err := c.BindJSON(&newCliente); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to obtain client"})
		return
	}

	if !services.AtualizarCliente(c_id, newCliente) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "client not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, newCliente)
}

func delCliente(c *gin.Context) {
	if ok, msg := checkToken(c, []string{"user", "gerent"}); !ok{ // medicos e atendentes nao podem apagar dados do cliente
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}
	id := c.Param("id")

	c_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "client id not valid"})
		return
	}
	if !services.DeletarCliente(c_id) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to delete client", "id": c_id})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "client deleted successfully", "id": c_id})
}

// Funcao para listar clientes
func getMedicos(c *gin.Context) {
	if ok, msg := checkToken(c, []string{}); !ok{ // qualquer um pode pesquisar medicos (passivel de alteracao)
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}
	var medicos = services.ListarMedicos()
	c.IndentedJSON(http.StatusOK, medicos)
}

// lista cliente pelo id
func getMedicosID(c *gin.Context) {
	if ok, msg := checkToken(c, []string{}); !ok{ // qualquer um pode pesquisar medicos por id
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}

	id := c.Param("id")
	m_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "medic id not valid"})
		return
	}

	medico, ok := services.LerMedico(m_id)
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "medic not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, medico)
}

func postMedico(c *gin.Context) {
	var newMedico services.Medico

	if ok, msg := checkToken(c, []string{"medico", "gerent"}); !ok{ // apenas medicos e gerentes podem enviar dados de medicos
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}

	if err := c.BindJSON(&newMedico); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to obtain medic"})
		return
	}

	if err := services.CriarMedico(newMedico); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newMedico)
}

func putMedico(c *gin.Context) {
	if ok, msg := checkToken(c, []string{"medico", "gerent"}); !ok{ // apenas medicos e gerentes podem alterar dados de medicos
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}

	id := c.Param("id")
	m_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "medic id not valid"})
		return
	}

	var newMedico services.Medico
	if err := c.BindJSON(&newMedico); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to obtain medic"})
		return
	}

	if !services.AtualizarMedico(m_id, newMedico) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to update medic", "id": m_id})
		return
	}
	c.IndentedJSON(http.StatusOK, newMedico)
}

func delMedico(c *gin.Context) {
	if ok, msg := checkToken(c, []string{"gerent"}); !ok{ // apenas gerentes podem deletar dados de medicos
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}
	id := c.Param("id")

	m_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "client id not valid"})
		return
	}
	if !services.DeletarMedico(m_id) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to delete medic", "id": m_id})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "medic deleted successfully", "id": m_id})
}

// funcoes para atendimentos
func postAtend(c *gin.Context) {
	if ok, msg := checkToken(c, []string{"user", "atende", "gerent"}); !ok{ // medicos nao podem solicitar atendimentos (passivel de alteracao)
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}
	var newAtend services.Atendimento
	if err := c.BindJSON(&newAtend); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to obtain apointment"})
		return
	}
	if err := services.SolicitarAtendimento(newAtend); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAtend)
}

// lista todos os atendimentos
func getAtends(c *gin.Context) {
	if ok, msg := checkToken(c, []string{"atende", "gerent"}); !ok{ // apenas atendentes e gerentes podem ler todos os atendimentos
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}

	var atendimentos = services.ListarAtendimentos()
	c.IndentedJSON(http.StatusOK, atendimentos)
}

// lista atendimentos por id de cliente
func getAtendsCliente(c *gin.Context) {
	if ok, msg := checkToken(c, []string{"user", "atende", "gerent"}); !ok{ // apenas medicos nao podem ver atendimentos por cliente
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}

	id := c.Param("id")
	c_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "client id not valid"})
		return
	}
	var atendimentos = services.LerAtendimentosCliente(c_id)
	c.IndentedJSON(http.StatusOK, atendimentos)
}

// lista atendimentos por id de medico
func getAtendsMedico(c *gin.Context) {
	if ok, msg := checkToken(c, []string{"medico", "atende", "gerent"}); !ok{ // apenas usuarios nao podem ver atendimentos por medico
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}

	id := c.Param("id")
	m_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "medic id not valid"})
		return
	}
	var atendimentos = services.LerAtendimentosMedico(m_id)
	c.IndentedJSON(http.StatusOK, atendimentos)
}

func putAtend(c *gin.Context) {
	if ok, msg := checkToken(c, []string{"atende", "gerent"}); !ok{ // apenas atendentes e gerentes podem alterar atendimentos
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}
	id := c.Param("id")
	a_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "appointment id not valid"})
		return
	}
	var newAtend services.Atendimento
	if err := c.BindJSON(&newAtend); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to obtain appointment"})
		return
	}

	if !services.AtualizarAtendimento(a_id, newAtend) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "appointment not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, newAtend)
}

func delAtend(c *gin.Context) {
	if ok, msg := checkToken(c, []string{"atende", "gerent"}); !ok{ // apenas atendentes e gerentes podem remover atendimentos
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}
	id := c.Param("id")

	a_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "appointment id not valid"})
		return
	}
	if !services.DeletarAtendimento(a_id) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to delete appointment", "id": a_id})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "appointment deleted successfully", "id": a_id})
}

// cria um documento
func postDoc(c *gin.Context) {
	if ok, msg := checkToken(c, []string{"medico", "gerent"}); !ok{ // apenas medicos e gerentes podem adicionar documentos
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}

	var newDoc services.Documento
	if err := c.BindJSON(&newDoc); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to obtain document"})
		return
	}

	if err := services.CriarDocumento(newDoc); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newDoc)
}

// lista todos os documentos
func getDocs(c *gin.Context) {
	if ok, msg := checkToken(c, []string{"medico", "gerent"}); !ok{ // apenas medicos e gerentes podem ver todos os documentos (passivel de alteracao)
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}

	var documentos = services.ListarDocumentos()
	c.IndentedJSON(http.StatusOK, documentos)
}

// lista documentos por id de cliente
func getDocsCliente(c *gin.Context) {
	if ok, msg := checkToken(c, []string{}); !ok{ // qualquer um pode ver documentos pelo cliente
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}
	id := c.Param("id")
	c_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "client id not valid"})
		return
	}
	var documentos = services.ListarDocumentosCliente(c_id)
	c.IndentedJSON(http.StatusOK, documentos)
}

// lista documentos por id de medico
func getDocsMedico(c *gin.Context) {
	if ok, msg := checkToken(c, []string{"atende", "medico", "gerent"}); !ok{ // apenas clientes nao podem ver documentos por medico
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}
	id := c.Param("id")
	m_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "medic id not valid"})
		return
	}
	var documentos = services.ListarDocumentosMedico(m_id)
	c.IndentedJSON(http.StatusOK, documentos)
}

func putDoc(c *gin.Context) {
	if ok, msg := checkToken(c, []string{"medico", "gerent"}); !ok{ // apenas medicos e gerentes podem alterar documentos
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}
	id := c.Param("id")
	d_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "appointment id not valid"})
		return
	}
	var newDoc services.Documento
	if err := c.BindJSON(&newDoc); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to obtain appointment"})
		return
	}

	if !services.AtualizarDocumento(d_id, newDoc) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "appointment not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, newDoc)
}

func delDoc(c *gin.Context) {
	if ok, msg := checkToken(c, []string{"medico", "gerent"}); !ok{ // apenas medicos e gerentes podem remover documentos
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}
	id := c.Param("id")

	d_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "doccument id not valid"})
		return
	}
	if !services.DeletarDocumento(d_id) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to delete doccument", "id": d_id})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "doccument deleted successfully", "id": d_id})
}

// funções para autenticação de usuário, usando pacote
// interno "auth" para operação, basicamente apenas
// extrai informações necessárias para serem enviadas
// ao pacote (24/06)

// struct para obter login e senha do usuario
type user_cred struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	ID       int    `json:"id"`
}

// Funcao de autenticacao, vulgo login
func postAuth(c *gin.Context) {
	var login user_cred
	// Usa BindJson para pegar o login dado pelo usuario
	if err := c.BindJSON(&login); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to obtain login credentials"})
		return
	}

	msg, ok := auth.LogUser(login.Username, login.Password)

	if !ok {
		// Se houver erro, mostra ao usuario qual o erro
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
		return
	}

	// Se estiver tudo ok, msg = token, envia ao usuario
	c.IndentedJSON(http.StatusOK, gin.H{"token": msg})
}

// Funcao de criacao de conta, vulgo sign up
func postAuthNew(c *gin.Context) {
	var login user_cred
	// Usa BindJson para pegar o login dado pelo usuario
	if err := c.BindJSON(&login); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to obtain login credentials"})
		return
	}

	if ok, _ := checkToken(c, []string{"gerent"}); ok { // apenas gerentes podem criar contas de outros cargos
		auth.CreateUser(login.Username, login.Password, login.Role, login.ID)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "user of role " + login.Role + " created successfully"})
		return
	}
	// se nao for gerente, cria uma conta de usuario independente do que o usuario digitou
	auth.CreateUser(login.Username, login.Password, "user", login.ID)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "user created successfully"})

}

// Faz a checagem to token | Alteração: converter role string em roles string[],
// pois tem funcoes que podem ser acessadas por mais de um cargo (24/06)
func checkToken(c *gin.Context, roles []string) (bool, string) {
	// Pega o token pelo header de autorizacao
	tokenString := c.GetHeader("Authorization")

	// Se nao tiver token, avisa que nao achou o token e retorna
	if tokenString == "" {
		return false, "token not found"
	}

	msg, ok := auth.CheckToken(tokenString, roles)

	return ok, msg
}
