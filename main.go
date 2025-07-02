package main

import (
	"net/http"

	"engsoft/auth"

	"github.com/gin-gonic/gin"

	"engsoft/services"
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
	//PUT
	router.PUT("/albums/:id", putAlbumID)
	//DELETE
	router.DELETE("/albums/:id", deleteAlbumID)

	//Funcoes da aplicacao final real oficial
	//GET
	router.GET("/clientes", getClientes)

	//POST
	router.POST("/clientes/novo", postCliente)
	router.POST("/atendimentos/novo", postAtend)
	//PUT

	//DELETE

	// roda o servidor localmente
	router.Run("localhost:8080")

	// OBS: Existe um meio de testar o token ao receber uma rota,
	// mas por agora (12/06) nao irei fazer isso
	// VVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVV
	// admin := router.group("/")
	// admin.Use(*Funcao para autenticar*){*rotas de admin aqui*}
}

// Funcoes para lidar com clientes
func getClientes(c *gin.Context) {
	var clientes = services.ListarClientes()
	c.IndentedJSON(http.StatusOK, clientes)
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

// funções para autenticação de usuário, usando pacote
// interno "auth" para operação, basicamente apenas
// extrai informações necessárias para serem enviadas
// ao pacote (24/06)

// struct para obter login e senha do usuario
type user_cred struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
