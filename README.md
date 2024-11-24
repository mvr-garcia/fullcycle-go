# Projeto de Cotação do Dólar

Este projeto é uma aplicação simples em Go que busca a cotação do dólar usando uma API externa e armazena os dados em um banco de dados SQLite. Ele demonstra o uso do pacote context para gerenciar timeouts em chamadas externas, operações de banco de dados e comunicação entre cliente e servidor.

## Objetivo

O objetivo deste projeto é implementar uma aplicação cliente-servidor que:

- Permita consultar a cotação atual do dólar a partir de uma API externa.
- Salve a cotação no banco de dados SQLite no servidor.
- Retorne apenas o valor da cotação (bid) para o cliente.
- Gere um arquivo cotacao.txt no cliente com o valor da cotação no formato:
    - `Dólar: {valor}`

### Instalação e Configuração
#### Pré-requisitos

Go instalado na sua máquina (versão 1.18 ou superior).
Conhecimento básico em linha de comando.
Acesso à internet para consultar a API externa.

#### Passo 1: Instalar o Go

Faça o download do Go para o seu sistema operacional a partir do site oficial: https://go.dev/dl/.
Siga as instruções de instalação fornecidas na página.

Para verificar se a instalação foi bem-sucedida, execute:

```shell
go version
```

Você verá algo como:

```shell
go version go1.20 darwin/amd64
```

Passo 2: Clonar o Repositório

```shell
git clone https://github.com/mvr-garcia/fullcycle-go
cd fullcycle-go
```

Passo 3: Executar o Servidor

Compile e execute o servidor:

```shell
go run server/main.go
```

O servidor será iniciado na porta 8080 e estará disponível no endpoint:
http://localhost:8080/cotacao

Passo 4: Executar o Cliente

Em outro terminal, compile e execute o cliente:

```shell
go run client/main.go
```

### Licença

Este projeto é fornecido "como está" para fins educacionais. Você é livre para usá-lo, modificá-lo e distribuí-lo conforme necessário. 🚀

Este README.md fornece todas as informações para qualquer pessoa instalar, rodar e entender o projeto! Se precisar de algo mais, avise! 😊
