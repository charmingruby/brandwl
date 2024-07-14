# Brand Monitor Challenge API

# Requisitos

- Go

# Execução

**1.** Adicionar o .env no diretório raiz com as configurações necessárias.

**2.** Baixar dependências

```
$ go mod tidy
```

**3.** Rodar projeto

```
$ go run cmd/api/main.go
```
## API

### Rotas

| Método | URL                                   | Privada | Funcionalidade                                                                            |
| :----- | ------------------------------------- | :-----: | ----------------------------------------------------------------------------------------- |
| POST   | /api/v1/search                         |    -    | Retorna o resultado da pesquisa por termo. |


### Dados para execução

#### `[POST]` /api/v1/search

Payload:

```
{
    "email": "gustavodiasa2121@gmail.com",
    "terms": [
        "test"
    ]
}
```
