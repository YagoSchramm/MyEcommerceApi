# MyEcommerce API

API REST para um sistema de e-commerce construída em Go.

## Visão Geral

Esta API oferece funcionalidades para gerenciamento de usuários, produtos, avaliações, compras e imagens.

## Requisitos

- Go 1.25.6 ou superior
- PostgreSQL
- Docker / Docker Compose

## Configuração de Ambiente

Copie o arquivo de exemplo para `.env`:

```powershell
copy .env-example .env
```

Edite `.env` caso queira alterar as credenciais ou porta.

## Executando o banco de dados com Docker

Para subir apenas o PostgreSQL:

```powershell
docker-compose up postgres
```

Para subir toda a stack (API + banco):

```powershell
docker-compose up
```

## Executando as migrations

Rode o comando abaixo para aplicar as migrations no banco:

```powershell
go run ./cmd/api -migrate
```

## Iniciando a API

Após o banco estar disponível e as migrations aplicadas:

```powershell
go run ./cmd/api
```

A API estará disponível em `http://localhost:8080`.

## Rotas principais

- `POST /users` - criar usuário (público)
- `POST /auth/login` - autenticar usuário (público)

### Rotas protegidas (token obrigatório)

- `POST /auth/logout` - encerrar sessão
- `POST /image/save` - enviar imagem
- `GET /uploads/{file}` - acessar arquivo de imagem
- `GET /users` - listar usuários
- `GET /users/{id}` - buscar usuário por ID
- `PUT /users/{id}` - atualizar usuário
- `DELETE /users/{id}` - excluir usuário
- `POST /products` - criar produto
- `GET /products` - listar produtos
- `GET /products/{id}` - buscar produto por ID
- `PUT /products/{id}` - atualizar produto
- `DELETE /products/{id}` - excluir produto
- `POST /ratings` - criar avaliação
- `GET /ratings/{id}` - buscar avaliação por ID
- `GET /ratings/user/{userId}` - avaliações por usuário
- `GET /ratings/product/{productId}` - avaliações por produto
- `PUT /ratings/{id}` - atualizar avaliação
- `DELETE /ratings/{id}` - excluir avaliação
- `POST /purchases` - criar compra
- `GET /purchases` - listar compras
- `GET /purchases/{id}` - buscar compra por ID
- `GET /purchases/user/{userId}` - compras por usuário

## Documentação de API

Veja `api.md` para o template Swagger/OpenAPI com todos os endpoints e rotas protegidas.
