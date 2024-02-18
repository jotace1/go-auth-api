<h1 align="center" style="font-weight: bold;">Go Auth API 💻</h1>

<p align="center">
 <a href="#tech">Technologies</a> • 
 <a href="#started">Getting Started</a> • 
  <a href="#routes">API Endpoints</a> •
 <a href="#colab">Collaborators</a> •
 <a href="#contribute">Contribute</a>
</p>

<p align="center">
    <b>Simple API with account creation and authentication</b>
</p>

<h2 id="technologies">💻 Technologies</h2>

- Golang
- Docker
- PostgreSQL
- Fiber

<h2 id="started">🚀 Getting started</h2>

<h3>Prerequisites</h3>

Here you list all prerequisites necessary for running your project. For example:

- [GoLang](https://go.dev/)
- [Docker](https://www.docker.com/)
- [Mockery](https://vektra.github.io/mockery/latest/)

<h3>Cloning</h3>

How to clone your project

```bash
git clone https://github.com/jotace1/go-auth-api
```

<h3>Config .env variables</h2>

Use the `.env.example` as reference to create your configuration file `.env` with your own Credentials

```yaml
DATABASE_URL={DATABASE_URL}
SIGNATURE_SECRET={JWT_SIGNATURE_SECRET}
DB_USER={DB_USER}
DB_PASSWORD={DB_PASSWORD}
DB_NAME={DB_NAME}
```

<h3>Starting</h3>

How to start your project

```bash
cd go-auth-api
mockery --all
docker-compose up
```

<h2 id="routes">📍 API Endpoints</h2>

Here you can list the main routes of your API, and what are their expected request bodies.
​
| route               | description                                          
|----------------------|-----------------------------------------------------
| <kbd>POST /account</kbd>     | creates an user account see [response details](#post-auth-account)
| <kbd>POST /account/login</kbd>     | authenticate user into the api see [request details](#post-auth-login)
| <kbd>POST /auth/test</kbd>     | test route to check if token is valid see [request details](#post-auth-route)

<h3 id="post-auth-account">POST /account</h3>

**REQUEST**
```json
{
  "username": "Joao",
  "email": "joao@email.com",
  "password": "1234567"
}
```


**RESPONSE**
```json
{
  "account_id": "5a37d698-78e1-470e-8eb7-94c5aa802b8e",
  "username": "Joao",
  "email": "joao@email.com"
}
```

<h3 id="post-auth-login">POST /account/login</h3>

**REQUEST**
```json
{
  "email": "joao123@email.com",
  "password": "1234567"
}
```

**RESPONSE**
```json
{
  "token": "OwoMRHsaQwyAgVoc3OXmL1JhMVUYXGGBbCTK0GBgiYitwQwjf0gVoBmkbuyy0pSi"
}
```

<h3 id="post-auth-route">POST /auth/test</h3>

**REQUEST HEADER**
```json
"Authorization": "Bearer OwoMRHsaQwyAgVoc3OXmL1JhMVUYXGGBbCTK0GBgiYitwQwjf0gVoBmkbuyy0pSi"
```

**RESPONSE**
```json
"You are authenticated"
```
