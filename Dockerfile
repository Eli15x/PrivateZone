# Use uma imagem base do Golang
FROM golang:latest

# Defina o diretório de trabalho dentro do container
WORKDIR /go/src/app

# Copie o código atual do diretório local para o diretório de trabalho no container
COPY . .

# Baixe as dependências do Go
RUN go mod download

# Comando padrão para executar o aplicativo quando o contêiner for iniciado
CMD ["go", "run", "main.go"]
