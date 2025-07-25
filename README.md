# Desenvolvendo API Rest em Go
Este projeto foi criado como parte do meu processo de aprendizado da linguagem Go (Golang). Ele consiste em uma API RESTful simples que realiza operações CRUD conectando-se a um banco de dados PostgreSQL via Docker, utilizando o framework gorilla/mux para o roteamento e gorm como ORM.

## 🚀 Funcionalidades

A API expõe os seguintes endpoints para gerenciar "personalidades":

GET /personalidades – Exibe todas as personalidades

GET /personalidades/{id} – Exibe uma personalidade por ID

POST /criarpersonalidades – Cria uma nova personalidade

DELETE /deletapersonalidades/{id} – Deleta uma personalidade

## 🛠️ Tecnologias Utilizadas
Go (Golang)

Gorilla Mux – Framework de roteamento

GORM – ORM para Go

PostgreSQL

Docker

## 🐳 Configuração com Docker
1. Clone o repositório:

```
bash

git clone https://github.com/andreylsant/Desenvolvendo_Api_Rest_golang.git
cd Desenvolvendo_Api_Rest_golang
```

2. Suba o container do PostgreSQL com Docker Compose:

**Observação:** Certifique-se de ter um ```docker-compose.yml```. Se ainda não tiver, posso te ajudar a montar.

```
bash

docker-compose up -d
```

3. Execute o projeto Go:
```
bash

go run main.go
```
A API estará rodando em: http://localhost:8000

## 🗃️ Estrutura do Projeto
````
Desenvolvendo_Api_Rest_golang/
├── controllers/        # Funções que controlam as requisições (CRUD de personalidades)
├── models/             # Estrutura dos dados (ex: Personalidade)
├── database/           # Conexão com o banco de dados
├── routes/             # Arquivo com as rotas da aplicação
├── main.go             # Ponto de entrada da aplicação
└── go.mod              # Gerenciador de dependências do Go
````

## 📌 Próximos passos (sugestões)

Adicionar autenticação JWT

Implementar atualização de dados (PUT)

Criar testes automatizados

Documentar a API com Swagger