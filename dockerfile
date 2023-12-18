# Use a imagem oficial do golang como base
FROM golang:latest

# Configure o diretório de trabalho dentro do contêiner
WORKDIR /go/src/simple-go-api/cmd/simple-go-api

# Copie o código fonte para o diretório de trabalho
COPY . .

# Baixe as dependências do Go
RUN go get -d -v ./...

# Compile o aplicativo
RUN go install -v ./...

# Exponha a porta que a aplicação vai ouvir
EXPOSE 8080

# Comando padrão para executar o aplicativo
CMD ["simple-go-api"]
