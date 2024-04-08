# Healthcheck Go

Este é um simples programa em Go que verifica o status de saúde de várias URLs listadas em um arquivo.

## Funcionalidades

- Lê uma lista de URLs de um arquivo.
- Faz uma requisição HTTP para cada URL e verifica o status da resposta.
- Registra o resultado (URL e status) em um arquivo de log.

## Pré-requisitos

- Go instalado em seu sistema.

## Como usar

1. Certifique-se de ter o Go instalado em sua máquina.
2. Clone este repositório: `git clone https://github.com/leonardofgc/healthcheck-go.git`
3. Navegue até o diretório do projeto: `cd healthcheck-go`
4. Crie um arquivo chamado `hosts.txt` contendo as URLs a serem verificadas, uma por linha.go`
5. Execute o programa: `go run main.
6. O resultado será registrado no arquivo `healthcheck`

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para enviar pull requests ou relatar problemas.

## Licença

Este projeto está licenciado sob a [Licença MIT](https://github.com/seu_usuario/healthcheck-go/blob/main/LICENSE).

