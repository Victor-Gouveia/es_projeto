package auth

// pacote utilizado para todas as operações relacionadas
// com autenticação de usuário, geração e checagem de token
import (
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
}

// lista de usuarios "hard-coded"
var users = []user{
	{Username: "test", Password: "test", Role: "user"},
	{Username: "admin", Password: "admin", Role: "admin"},
}

func GenerateToken(username string, role string) string {
	// Faz os claims pelo jwt para criar o token com eles
	// OBS: Idealmente teria a data de expiracao do token,
	// mas nao coloquei para simplificar implementação e
	// testes
	claims := jwt.MapClaims{
		"username": username,
		"role":     role,
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

func LogUser(username string, password string) (string, bool) {

	// Loop para procurar usuario e senha corretos
	for _, a := range users {
		if a.Username == username && a.Password == password {
			// Obtem o cargo do usuario para gerar o token
			var role = a.Role
			// Gera o token a partir do usuario dado
			sign_token := GenerateToken(username, role)

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
		if !ok {
			// Se nao for valido o cargo, retorna o erro ao usuario
			return "role sign not valid", false
		}
		// Se precisar de cargo e nao tiver, retorna que nao esta autorizado
		if slices.Contains(roles, tknRole) && len(roles) > 0 {
			return "token does not have required role", false
		}
		// Estando tudo ok, retorna que a checagem deu certo
		return "", true
	}
	// Erro inesperado, esta aqui de fallback, nao deve ocorrer
	return "there was an issue validating your token", false
}
