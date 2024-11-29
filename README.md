# Go Expert - Stress Test

Implementar uma app em go usando uma CLI para fazer uma quantidade determinada de chamadas para uma URL.

## Como executar

### Via Docker

O projeto está disponível no Docker Hub, e para executar o container basta utilizar o comando abaixo, substituindo os valores de `--url`, `--requests` e `--concurrency` pelos valores desejados.

```sh
docker run fahds1993/stress-test:latest \
    --url https://google.com.br \
    --requests 100 \
    --concurrency 10
```

### Localmente

Para executar o projeto localmente, é necessário ter o Go instalado na máquina. Após a instalação, basta executar o comando abaixo, substituindo os valores de `--url`, `--requests` e `--concurrency` pelos valores desejados.

```sh
go run cmd/cli/main.go \
    --url https://google.com.br \
    --requests 100 \
    --concurrency 10
```

## Testes

Para executar os testes de unidade, basta executar o comando abaixo.

```sh
make test
```
