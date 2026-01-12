# 🍽️ Mamajulia API

## 🎯 Sobre o Projeto

API desenvolvida para gerenciar um restaurante com as seguintes funcionalidades:

- ✅ Sistema de autenticação (Signup/Signin)
- ✅ Controle de acesso por roles (user/admin)
- ✅ CRUD de pratos (apenas admin pode criar/editar/deletar)
- ✅ Visualização pública de pratos
- ✅ Sistema de pedidos com status (em_andamento/entregue)
- ✅ Middleware de autenticação e autorização

## 🛠️ Tecnologias

- **Go 1.25+**
- **Gin**
- **GORM**
- **Neon - PostgreSQL**
- **JWT**
- **bcrypt**

3. Configure as variáveis de ambiente criando um arquivo `.env` na raiz do projeto:
```env
DATABASE_URL=postgresql://usuario:senha@host:porta/banco?sslmode=require
JWT_SECRET=123u81ue01j0rf21n3
```

4. Execute a aplicação:
```bash
go run cmd/api/main.go
```

A API estará rodando em `http://localhost:3000`

## 🔐 Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

| Variável | Descrição | Obrigatório |
|----------|-----------|-------------|
| `DATABASE_URL` | String de conexão PostgreSQL (formato: `postgresql://user:password@host:port/dbname?sslmode=require`) | Sim |
| `JWT_SECRET` | Chave secreta para assinar tokens JWT | Não (usa "default" se não definido) |

**Exemplo de `.env`:**
```env
DATABASE_URL=postgresql://user:password@ep-xxx-xxx.us-east-2.aws.neon.tech/neondb?sslmode=require
JWT_SECRET=minha_chave_secreta_super_segura_12345
```

## 📚 Estrutura de Dados

### User (Usuário)
```json
{
  "id": 1,
  "name": "",
  "email": "",
  "role": "user" // ou "adm" para administrador
}
```

### Dish (Prato)
```json
{
  "id": 1,
  "name": "Pizza Margherita",
  "description": "Pizza tradicional italiana",
  "value": 35.90,
  "image": "https://exemplo.com/imagem.jpg"
}
```

### Order (Pedido)
```json
{
  "id": 1,
  "mesa_number": 5,
  "status": "em_andamento", // ou "entregue"
  "user_id": 1,
  "user": {
    "id": 1,
    "name": "João Silva",
    "email": "joao@email.com"
  },
  "dishes": [
    {
      "id": 1,
      "dish_id": 1,
      "quantity": 2,
      "dish": {
        "id": 1,
        "name": "Pizza Margherita",
        "description": "Pizza tradicional italiana",
        "value": 35.90
      }
    }
  ],
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

### Como obter o token:
1. Faça o cadastro em `POST /auth/signup`
2. Faça login em `POST /auth/signin`
3. Use o token retornado no header `Authorization` das requisições protegidas

**Token expira em 24 horas**

## 📡 Rotas da API

### 🔓 Autenticação (Públicas)

#### POST /auth/signup
Cadastra um novo usuário.

**Request Body:**
```json
{
  "name": "João Silva",
  "email": "joao@email.com",
  "password": "senha123",
  "role": "user" // opcional, padrão: "user"
}
```

#### POST /auth/signin
Faz login e retorna o token JWT.

**Request Body:**
```json
{
  "email": "joao@email.com",
  "password": "senha123"
}
```

### 🍽️ Pratos (Públicas)

#### GET /pratos
Lista todos os pratos disponíveis.

#### GET /pratos/:id
Busca um prato específico por ID.

### 🔒 Pedidos (Requer Autenticação)

#### POST /pedidos
Cria um novo pedido. O usuário é identificado automaticamente pelo token JWT.


### 👑 Administração (Requer Autenticação + Role "adm")

#### POST /admin/pratos
Cria um novo prato.

#### PUT /admin/pratos/:id
Atualiza um prato existente.

#### DELETE /admin/pratos/:id
Deleta um prato.

#### GET /admin/pedidos
Lista todos os pedidos (apenas admin).

#### GET /admin/pedidos/:id
Busca um pedido específico por ID (apenas admin).

#### PUT /admin/pedidos/:id/status
Atualiza o status de um pedido (apenas admin).

#### DELETE /admin/pedidos/:id
Deleta um pedido (apenas admin).


## 📊 Resumo das Rotas

| Método | Rota | Autenticação | Role | Descrição |
|--------|------|--------------|------|-----------|
| POST | `/auth/signup` | ❌ | - | Cadastra novo usuário |
| POST | `/auth/signin` | ❌ | - | Faz login e retorna token |
| GET | `/pratos` | ❌ | - | Lista todos os pratos |
| GET | `/pratos/:id` | ❌ | - | Busca prato por ID |
| POST | `/pedidos` | ✅ | user/admin | Cria novo pedido |
| POST | `/admin/pratos` | ✅ | **adm** | Cria novo prato |
| PUT | `/admin/pratos/:id` | ✅ | **adm** | Atualiza prato |
| DELETE | `/admin/pratos/:id` | ✅ | **adm** | Deleta prato |
| GET | `/admin/pedidos` | ✅ | **adm** | Lista todos os pedidos |
| GET | `/admin/pedidos/:id` | ✅ | **adm** | Busca pedido por ID |
| PUT | `/admin/pedidos/:id/status` | ✅ | **adm** | Atualiza status do pedido |
| DELETE | `/admin/pedidos/:id` | ✅ | **adm** | Deleta pedido |

## 🔒 Níveis de Acesso

- **Público**: Qualquer pessoa pode acessar
- **Autenticado**: Requer token JWT válido (usuário ou admin)
- **Admin**: Requer token JWT válido + role "adm"

## 📝 Validações Importantes

### Usuário (Signup)
- `name`: obrigatório
- `email`: obrigatório, formato válido de email
- `password`: obrigatório, mínimo 6 caracteres
- `role`: opcional, padrão "user" (valores: "user" ou "adm")

### Prato
- `name`: obrigatório
- `description`: obrigatório
- `value`: obrigatório, deve ser maior que 0
- `image`: opcional (gera placeholder automático se vazio)

### Pedido
- `mesa_number`: obrigatório, número inteiro
- `dishes`: obrigatório, array não vazio
  - `dish_id`: obrigatório, ID do prato existente
  - `quantity`: obrigatório, número maior que 0
- `status`: padrão "em_andamento" (valores válidos: "em_andamento", "entregue")

## 🚨 Códigos de Status HTTP

| Código | Significado |
|--------|-------------|
| 200 | OK - Requisição bem-sucedida |
| 201 | Created - Recurso criado com sucesso |
| 400 | Bad Request - Dados inválidos |
| 401 | Unauthorized - Token não fornecido ou inválido |
| 403 | Forbidden - Acesso negado (sem permissão) |
| 404 | Not Found - Recurso não encontrado |
| 500 | Internal Server Error - Erro no servidor |

## 💡 Exemplos de Uso

### 1. Criar um usuário admin
```bash
curl -X POST http://localhost:3000/auth/signup \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Admin",
    "email": "admin@email.com",
    "password": "admin123",
    "role": "adm"
  }'
```

### 2. Fazer login
```bash
curl -X POST http://localhost:3000/auth/signin \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@email.com",
    "password": "admin123"
  }'
```

### 3. Criar um prato (como admin)
```bash
curl -X POST http://localhost:3000/admin/pratos \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <seu_token>" \
  -d '{
    "name": "Pizza Margherita",
    "description": "Pizza tradicional italiana",
    "value": 35.90,
    "image": "https://exemplo.com/pizza.jpg"
  }'
```

### 4. Listar pratos
```bash
curl -X GET http://localhost:3000/pratos
```

### 5. Criar um pedido
```bash
curl -X POST http://localhost:3000/pedidos \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <seu_token>" \
  -d '{
    "mesa_number": 5,
    "dishes": [
      {
        "dish_id": 1,
        "quantity": 2
      }
    ]
  }'
```

## 🗄️ Banco de Dados

A API utiliza GORM para migração automática das tabelas:
- `users` - Usuários do sistema
- `dishes` - Pratos do cardápio
- `orders` - Pedidos
- `order_dishes` - Itens dos pedidos (relação many-to-many)

As migrações são executadas automaticamente ao iniciar a aplicação.
