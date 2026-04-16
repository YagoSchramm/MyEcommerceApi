# API Documentation

This document provides a Swagger/OpenAPI-style template for the MyEcommerce API.

## Overview

- Base URL: `http://localhost:8080`
- Authentication: JWT Bearer token in `Authorization: Bearer <token>`
- Protected routes require the `Authorization` header.

```yaml
openapi: 3.0.3
info:
  title: MyEcommerce API
  version: 1.0.0
  description: API REST para o sistema de e-commerce MyEcommerce.
servers:
  - url: http://localhost:8080
components:
  securitySchemes:
    bearerAuth:
      type: http
      scheme: bearer
      bearerFormat: JWT
  schemas:
    UserCreate:
      type: object
      required:
        - name
        - email
        - password
        - roles
      properties:
        name:
          type: string
        email:
          type: string
          format: email
        password:
          type: string
          minLength: 6
        roles:
          type: array
          items:
            type: string
            enum: [buyer, seller, admin]
    UserResponse:
      type: object
      properties:
        id:
          type: string
          format: uuid
        name:
          type: string
        email:
          type: string
          format: email
        roles:
          type: array
          items:
            type: string
            enum: [buyer, seller, admin]
    LoginRequest:
      type: object
      required:
        - email
        - password
      properties:
        email:
          type: string
          format: email
        password:
          type: string
    LoginResponse:
      type: object
      properties:
        access_token:
          type: string
        refresh_token:
          type: string
    ProductCreate:
      type: object
      required:
        - user_id
        - user_name
        - name
        - value
        - image
        - stock
        - description
      properties:
        user_id:
          type: string
          format: uuid
        user_name:
          type: string
        name:
          type: string
        value:
          type: number
          format: float
        image:
          type: string
        stock:
          type: integer
        description:
          type: string
    RatingCreate:
      type: object
      required:
        - user_id
        - user_name
        - purchase_id
        - product_id
        - rating
      properties:
        user_id:
          type: string
          format: uuid
        user_name:
          type: string
        purchase_id:
          type: string
          format: uuid
        product_id:
          type: string
          format: uuid
        rating:
          type: number
          format: float
          minimum: 0
          maximum: 5
        description:
          type: string
    PurchaseCreate:
      type: object
      required:
        - product_id
        - user_id
        - value
        - quantity
      properties:
        product_id:
          type: string
          format: uuid
        user_id:
          type: string
          format: uuid
        value:
          type: number
          format: float
        quantity:
          type: integer
paths:
  /users:
    post:
      summary: Criar usuário
      description: Cria um usuário novo.
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/UserCreate'
      responses:
        '201':
          description: Usuário criado com sucesso
        '400':
          description: Dados inválidos
    get:
      summary: Listar usuários
      security:
        - bearerAuth: []
      responses:
        '200':
          description: Lista de usuários
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/UserResponse'
        '401':
          description: Não autorizado
  /users/{id}:
    get:
      summary: Buscar usuário por ID
      security:
        - bearerAuth: []
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
            format: uuid
      responses:
        '200':
          description: Usuário encontrado
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/UserResponse'
        '401':
          description: Não autorizado
    put:
      summary: Atualizar usuário
      security:
        - bearerAuth: []
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
            format: uuid
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              properties:
                name:
                  type: string
                roles:
                  type: array
                  items:
                    type: string
                    enum: [buyer, seller, admin]
      responses:
        '200':
          description: Usuário atualizado
        '400':
          description: Requisição inválida
        '401':
          description: Não autorizado
    delete:
      summary: Excluir usuário
      security:
        - bearerAuth: []
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
            format: uuid
      responses:
        '200':
          description: Usuário excluído
        '401':
          description: Não autorizado
  /auth/login:
    post:
      summary: Login
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/LoginRequest'
      responses:
        '200':
          description: Autenticação bem-sucedida
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/LoginResponse'
        '401':
          description: Credenciais inválidas
  /auth/logout:
    post:
      summary: Logout
      description: Encerra a sessão. Atualmente não exige autenticação.
      responses:
        '200':
          description: Logout bem-sucedido
  /image/save:
    post:
      summary: Enviar imagem
      description: Faz upload de imagem multipart/form-data.
      requestBody:
        required: true
        content:
          multipart/form-data:
            schema:
              type: object
              properties:
                image:
                  type: string
                  format: binary
      responses:
        '200':
          description: URL da imagem salva
  /products:
    post:
      summary: Criar produto
      security:
        - bearerAuth: []
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/ProductCreate'
      responses:
        '201':
          description: Produto criado
        '401':
          description: Não autorizado
    get:
      summary: Listar produtos
      security:
        - bearerAuth: []
      responses:
        '200':
          description: Lista de produtos
  /products/{id}:
    get:
      summary: Buscar produto por ID
      security:
        - bearerAuth: []
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
            format: uuid
      responses:
        '200':
          description: Produto encontrado
    put:
      summary: Atualizar produto
      security:
        - bearerAuth: []
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
            format: uuid
      responses:
        '200':
          description: Produto atualizado
    delete:
      summary: Excluir produto
      security:
        - bearerAuth: []
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
            format: uuid
      responses:
        '200':
          description: Produto excluído
  /ratings:
    post:
      summary: Criar avaliação
      security:
        - bearerAuth: []
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/RatingCreate'
      responses:
        '201':
          description: Avaliação criada
  /ratings/{id}:
    get:
      summary: Buscar avaliação por ID
      security:
        - bearerAuth: []
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
            format: uuid
      responses:
        '200':
          description: Avaliação encontrada
    put:
      summary: Atualizar avaliação
      security:
        - bearerAuth: []
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
            format: uuid
      responses:
        '200':
          description: Avaliação atualizada
    delete:
      summary: Excluir avaliação
      security:
        - bearerAuth: []
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
            format: uuid
      responses:
        '200':
          description: Avaliação excluída
  /ratings/user/{userId}:
    get:
      summary: Avaliações de um usuário
      security:
        - bearerAuth: []
      parameters:
        - name: userId
          in: path
          required: true
          schema:
            type: string
            format: uuid
      responses:
        '200':
          description: Lista de avaliações do usuário
  /ratings/product/{productId}:
    get:
      summary: Avaliações de um produto
      security:
        - bearerAuth: []
      parameters:
        - name: productId
          in: path
          required: true
          schema:
            type: string
            format: uuid
      responses:
        '200':
          description: Lista de avaliações do produto
  /purchases:
    post:
      summary: Criar compra
      security:
        - bearerAuth: []
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/PurchaseCreate'
      responses:
        '201':
          description: Compra criada
    get:
      summary: Listar compras
      security:
        - bearerAuth: []
      responses:
        '200':
          description: Lista de compras
  /purchases/{id}:
    get:
      summary: Buscar compra por ID
      security:
        - bearerAuth: []
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
            format: uuid
      responses:
        '200':
          description: Compra encontrada
  /purchases/user/{userId}:
    get:
      summary: Compras por usuário
      security:
        - bearerAuth: []
      parameters:
        - name: userId
          in: path
          required: true
          schema:
            type: string
            format: uuid
      responses:
        '200':
          description: Lista de compras do usuário
```
