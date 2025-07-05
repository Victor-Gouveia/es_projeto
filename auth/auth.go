package auth

// pacote utilizado para todas as operações relacionadas
// com autenticação de usuário, geração e checagem de token
import (
	//"fmt" // usado para ver alguns erros no console
	"slices"

	"github.com/golang-jwt/jwt/v5"
)

// chave secreta para criptografia do token
var secret = "chave_secreta"

// struct com login, senha e cargo do usuario
type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	ID       int    `json:"id"`
}

// lista de usuarios "hard-coded"
var users = []user{
	{Username: "test", Password: "test", Role: "user", ID: 1}, // "João Silva"
	{Username: "atende", Password: "atende", Role: "atende"},
	{Username: "medico", Password: "medico", Role: "medico", ID: 92}, // "Dra. Lucia Ferreira"
	{Username: "gerent", Password: "gerent", Role: "gerent"},
}

func GenerateToken(username string, role string, id int) string {
	// Faz os claims pelo jwt para criar o token com eles
	// OBS: Idealmente teria a data de expiracao do token,
	// mas nao coloquei para simplificar implementação e
	// testes
	claims := jwt.MapClaims{
		"username": username,
		"role":     role,
		"id":       id,
	}

	// Gera o token a partir dos claims, assina em HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret)) // assina com o segredo

	// Se houver erro, retorna uma string vazia, isso e
	// lidado depois em LogUser
	if err != nil {
		return ""
	}
	return tokenString
}

var cert_roles = []string{"user", "atende", "medico", "gerent"} // lista de cargos para conferir ao criar novo usuario

func CreateUser(username string, password string, role string, id int) bool {
	// conferindo se o cargo esta correto
	if !slices.Contains(cert_roles, role) {
		return false
	}
	var newUser user = user{Username: username, Password: password, Role: role, ID: id}

	users = append(users, newUser)
	return true
}

func LogUser(username string, password string) (string, bool) {
	// Loop para procurar usuario e senha corretos
	for _, a := range users {
		if a.Username == username && a.Password == password {
			// Obtem o cargo do usuario para gerar o token
			var role = a.Role
			var id = a.ID
			// Gera o token a partir do usuario dado
			sign_token := GenerateToken(username, role, id)

			// Se o token nao for nulo, retorna ao usuario
			if sign_token != "" {
				return sign_token, true
			}
			// Se o token for nulo, exibe erro
			return "error generating token", false
		}
	}
	// Se nao achar nenhum usuario com senha, mostra erro pro usuario
	return "user not found or wrong password", false
}

// usando array "roles" pois havera acoes que podem ter permissao de mais de um tipo de usuario (exemplo: )
func CheckToken(tokenString string, roles []string) (string, bool) {
	// Usa jwt.Parse para obter as informacoes do token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// Se nao estiver codificado no metodo certo, retorna um erro
			return nil, jwt.ErrSignatureInvalid
		}
		// Se tiver ok, retorna os claims
		return []byte(secret), nil
	})

	// Em caso de erro, retorna token invalido
	if err != nil {
		return "invalid token", false
	}

	// Obtem os claims do token se ele for valido
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Primeiro, testa se a informacao de admin esta correta
		tknRole, ok := claims["role"].(string)
		if !ok { // confere se houve erro ao obter cargo ou se nao eh um cargo valido
			// Se nao for valido o cargo, retorna o erro ao usuario
			return "unauthorized: token role sign not valid", false
		}
		// Se precisar de cargo e nao tiver, retorna que nao esta autorizado
		if !slices.Contains(roles, tknRole) && len(roles) > 0 {
			return "unauthorized: token does not have required role", false
		}
		// Estando tudo ok, retorna que a checagem deu certo
		return "", true
	}
	// Erro inesperado, esta aqui de fallback, nao deve ocorrer
	return "unauthorized: there was an issue validating your token", false
}
