package main

import (
	"net/http"

	"engsoft/auth"

	"github.com/gin-gonic/gin"

	"engsoft/services"

	"strconv"
)

// struct com informacoes dos albums, tudo com parametros json
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	router := gin.Default()
	//GET
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	//POST
	router.POST("/albums", postAlbums)
	router.POST("/auth", postAuth)
	router.POST("/auth/new", postAuthNew)
	//PUT
	router.PUT("/albums/:id", putAlbumID)
	//DELETE
	router.DELETE("/albums/:id", deleteAlbumID)

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
	//router.PUT("/documentos/:id", putDoc)
	//DELETE
	router.DELETE("/clientes/:id", delCliente)
	router.DELETE("/medicos/:id", delMedico)
	router.DELETE("/atendimentos/:id", delAtend)
	//router.DELETE("/documentos/:id", delDoc)

	// roda o servidor localmente
	router.Run("localhost:8080")

	// OBS: Existe um meio de testar o token ao receber uma rota,
	// mas por agora (12/06) nao irei fazer isso
	// VVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVV
	// admin := router.group("/")
	// admin.Use(*Funcao para autenticar*){*rotas de admin aqui*}
}

// Funcao para listar clientes
func getClientes(c *gin.Context) {
	var clientes = services.ListarClientes()
	c.IndentedJSON(http.StatusOK, clientes)
}

// lista cliente pelo id
func getClientesID(c *gin.Context) {
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
	var medicos = services.ListarMedicos()
	c.IndentedJSON(http.StatusOK, medicos)
}

// lista cliente pelo id
func getMedicosID(c *gin.Context) {
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
	var atendimentos = services.ListarAtendimentos()
	c.IndentedJSON(http.StatusOK, atendimentos)
}

// lista atendimentos por id de cliente
func getAtendsCliente(c *gin.Context) {
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
	var documentos = services.ListarDocumentos()
	c.IndentedJSON(http.StatusOK, documentos)
}

// lista documentos por id de cliente
func getDocsCliente(c *gin.Context) {
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
	id := c.Param("id")
	m_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "medic id not valid"})
		return
	}
	var documentos = services.ListarDocumentosMedico(m_id)
	c.IndentedJSON(http.StatusOK, documentos)
}

// ===========================================================
// ===== CODIGO ANTIGO - TUTORIAL GOLANG - APAGAR DEPOIS =====
// ===========================================================
// albums definidos "hard-coded"
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	// Faz o teste do token | Nao precisa de acesso de administrador
	if !checkToken(c, []string{}) {
		return
	}
	// Retorna a lista de albums
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID - Acha um album pelo ID dado e exibe
func getAlbumByID(c *gin.Context) {
	// Faz o teste do token | Nao precisa de acesso de administrador
	if !checkToken(c, []string{}) {
		return
	}

	id := c.Param("id")

	// Loop para achar um album com o mesmo ID
	for _, a := range albums {
		if a.ID == id {
			// Retorna o album achado ao usuario
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	// Se nao achar um album, retorna uma mensagem de erro
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	// Faz o teste do token | Precisa de acesso de administrador
	if !checkToken(c, []string{"admin"}) {
		return
	}

	var newAlbum album

	// Usa BindJson para pegar o album novo
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Adiciona o album novo na lista
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func putAlbumID(c *gin.Context) {
	// Faz o teste do token | Precisa de acesso de administrador
	if !checkToken(c, []string{"admin"}) {
		return
	}

	id := c.Param("id")

	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Loop para achar um album com o mesmo ID
	for i := 0; i < len(albums); i++ {
		if albums[i].ID == id {
			// Achando o album, testa as entradas e se nao forem nulas
			// substitui os valores
			//
			// Se nao testar antes ele coloca a string vazia
			if newAlbum.Artist != "" {
				albums[i].Artist = newAlbum.Artist
			}
			if newAlbum.Title != "" {
				albums[i].Title = newAlbum.Title
			}
			if newAlbum.Price > 0 {
				albums[i].Price = newAlbum.Price
			}

			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}
	// Se nao achar album, retorna um erro ao usuario
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbumID(c *gin.Context) {
	// Faz o teste do token | Precisa de acesso de administrador
	if !checkToken(c, []string{"admin"}) {
		return
	}

	id := c.Param("id")

	// Loop para achar um album com o mesmo ID
	for i := 0; i < len(albums); i++ {
		if albums[i].ID == id {
			// Golang nao tem splice, entao faz um append entre tudo
			// que ta antes do item achado com o que tiver depois
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "removed album", "id": id})
			return
		}
	}
	// Se nao achar album, retorna um erro ao usuario
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// ==================================================
// ===== FIM DO CODIGO ANTIGO - TUTORIAL GOLANG =====
// ==================================================

// funções para autenticação de usuário, usando pacote
// interno "auth" para operação, basicamente apenas
// extrai informações necessárias para serem enviadas
// ao pacote (24/06)

// struct para obter login e senha do usuario
type user_cred struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
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

	auth.CreateUser(login.Username, login.Password, login.Role)
}

// Faz a checagem to token | Alteração: converter role string em roles string[],
// pois tem funcoes que podem ser acessadas por mais de um cargo (24/06)
func checkToken(c *gin.Context, roles []string) bool {
	// Pega o token pelo header de autorizacao
	tokenString := c.GetHeader("Authorization")

	// Se nao tiver token, avisa que nao achou o token e retorna
	if tokenString == "" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "token not found"})
		return false
	}

	msg, ok := auth.CheckToken(tokenString, roles)

	if !ok {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": msg})
	}

	return ok
}
