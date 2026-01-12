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


