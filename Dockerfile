# Imagem oficial do Golang
FROM golang:1.22

# Diretório de trabalho dentro do container
WORKDIR /app

# Copiando o arquivo go mod e sum
COPY go.mod go.sum ./

# Baixando todas as dependências
RUN go mod download

# Copiando o código-fonte para o container
COPY . .

# Compilando a aplicação
RUN go build -o main .

# Expondo a porta 8000
EXPOSE 8000

# Comando para iniciar a aplicação
CMD ["./main"]