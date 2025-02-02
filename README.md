

# API GoLangEstudos

Esta é uma API desenvolvida em Go (Golang) para estudos, com funcionalidades relacionadas a usuários, publicações, autenticação e interações sociais, como seguir usuários e curtir publicações.

## 📋 Requisitos

- Go 1.23.5 ou superior.
- MySQL (ou outro banco de dados compatível com `go-sql-driver/mysql`).
- Git (para clonar o repositório).

## 🚀 Como executar o projeto

### 1. Clonar o repositório

```bash
git clone https://github.com/seu-usuario/seu-repositorio.git
cd seu-repositorio
```

### 2. Configurar o ambiente

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis de ambiente:

```env
DB_USUARIO=golang
DB_SENHA=golang
DB_NOME=devbook
API_PORT=5000
SECRET_KEY=Ma/YbSgmMe8vw/9Lh2dbgfIiGikW6QnkG5mXewBLhfsatVAAZf+X5T3R30TewQqPsKkyYAAxRlLCDFTJlnealw==
```

### 3. Configurar o banco de dados

Execute o script SQL abaixo para criar o banco de dados e as tabelas necessárias:

```sql
CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(120) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE seguidores(
    usuario_id int not null,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,
    seguidor_id int not null,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,
    primary key(usuario_id, seguidor_id)
) ENGINE=INNODB;

CREATE TABLE publicacoes(
    id int auto_increment primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not null,
    autor_id int not null,
    FOREIGN KEY (autor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,
    curtidas int default 0,
    criadaEm timestamp default current_timestamp
) ENGINE=INNODB;
```

### 4. Instalar dependências

Certifique-se de que todas as dependências do projeto estão instaladas. Execute:

```bash
go mod tidy
```

### 5. Executar a API

Para iniciar o servidor da API, execute:

```bash
go run main.go
```

A API estará disponível em `http://localhost:5000`.

---

## 📚 Endpoints da API

Aqui estão os principais endpoints disponíveis na API:

### Usuários

- **Criar usuário**: `POST /usuarios`
  - Corpo da requisição:
    ```json
    {
      "nome": "Usuario 1",
      "email": "usuario1@gmail.com",
      "nick": "usuario1",
      "senha": "123"
    }
    ```

- **Login**: `POST /login`
  - Corpo da requisição:
    ```json
    {
      "email": "usuario1@gmail.com",
      "senha": "123"
    }
    ```

- **Buscar usuários**: `GET /usuarios?usuario=teste`
- **Buscar usuário por ID**: `GET /usuarios/{id}`
- **Atualizar usuário**: `PUT /usuarios/{id}`
- **Atualizar senha do usuário**: `PUT /usuarios/{id}`
- **Deletar usuário**: `DELETE /usuarios/{id}`

### Seguidores

- **Seguir usuário**: `POST /usuarios/{id}/seguir`
- **Parar de seguir usuário**: `POST /usuarios/{id}/seguir`
- **Buscar seguidores de um usuário**: `GET /usuarios/{id}/seguidores`
- **Buscar usuários que um usuário segue**: `GET /usuarios/{id}/seguindo`

### Publicações

- **Criar publicação**: `POST /publicacoes`
  - Corpo da requisição:
    ```json
    {
      "titulo": "Minha publicação",
      "conteudo": "Conteúdo da publicação"
    }
    ```

- **Buscar publicação por ID**: `GET /publicacoes/{id}`
- **Buscar todas as publicações**: `GET /publicacoes`
- **Atualizar publicação**: `PUT /publicacoes/{id}`
- **Deletar publicação**: `DELETE /publicacoes/{id}`
- **Buscar publicações de um usuário**: `GET /usuarios/{id}/publicacoes`

### Curtidas

- **Curtir publicação**: `POST /publicacoes/{id}/curtir`
- **Descurtir publicação**: `POST /publicacoes/{id}/descurtir`

---

## 🛠️ Bibliotecas utilizadas

- **[Gorilla Mux](https://github.com/gorilla/mux)**: Para roteamento HTTP.
- **[JWT Go](https://github.com/dgrijalva/jwt-go)**: Para autenticação via tokens JWT.
- **[Go MySQL Driver](https://github.com/go-sql-driver/mysql)**: Para conexão com o banco de dados MySQL.
- **[Godotenv](https://github.com/joho/godotenv)**: Para gerenciamento de variáveis de ambiente.
- **[Badoux Checkmail](https://github.com/badoux/checkmail)**: Para validação de e-mails.
- **[Golang Crypto](https://pkg.go.dev/golang.org/x/crypto)**: Para criptografia de senhas.

