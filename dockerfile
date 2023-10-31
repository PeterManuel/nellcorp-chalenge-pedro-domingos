# Use a imagem oficial do Ubuntu como base
FROM golang:latest

WORKDIR /app

# Copie o código fonte da sua aplicação para o diretório de trabalho
COPY mybankapi/ .

# Instale a biblioteca Gorilla Mux
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/lib/pq

RUN go install github.com/gorilla/mux
RUN go install github.com/lib/pq
RUN go mod tidy

# Compile a aplicação Go
RUN go build -o mybankapi

# Defina o comando de inicialização da aplicação
CMD ["./mybankapi"]