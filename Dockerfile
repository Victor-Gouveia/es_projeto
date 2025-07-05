# Define golang 1.22 como builder
FROM golang:1.24-alpine AS builder

# Define o diretório de trabalho do container
WORKDIR /app

# Copia os arquivos de gerenciamento de dependências primeiro.
COPY go.mod go.sum ./

# Baixa as dependências do projeto
RUN go mod download

# Copia todo o resto do código-fonte do projeto para o container
COPY . .

# Compila a aplicação Go.
# - CGO_ENABLED=0: Cria um binário estaticamente linkado, sem dependências de bibliotecas C do sistema.
# -o /app/main: Especifica que o arquivo de saída será chamado 'main' e estará no diretório '/app'.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main .

# Decidiu usar 'alpine' por ser uma imagem base mínima. Versão fixada para não causar conflitos
FROM alpine:3.22.0

# Instala o curl para testes no container
RUN apk --no-cache add curl

# Define o diretório de trabalho
WORKDIR /app

# Copia APENAS o binário compilado do estágio 'builder' para a nossa imagem final.
# Sem incluir código-fonte ou ferramenta de compilação.
COPY --from=builder /app/main .

# Expõe a porta 8080.
EXPOSE 8080

# Iniciar a aplicação quando o container for executado.
CMD ["./main"]
