# Simple Clinic

![Build Status](https://api.travis-ci.com/RianAlvesAce/gestao_online.svg?branch=main)

O Simple Clinic é um sistema de gestão e análise de desempenho para clínicas médicas, com o intuito de facilitar a marcação de consultas, gerenciamento de despesas, gerenciamento de pacientes e com possibilidade de melhorias futuras.

> "Tudo deveria se tornar o mais simples possível, mas não simplificado." 
>
> \- Albert Einstein

## Funcionalidades

- FeedBack visual sobre os números referente a clinica (quantidade de pacientes atendidos, ganhos totais, forma de pagamentos usados) com filtro de periodos.
- Registro dos pacientes atendidos pela clínica
- Agendamento prévio de consultas no sistema
- Registro de pagamento, caso não possua convênio médico
- Coming soon ✨

## Pré-requisitos para usar o sistema

- Possuir o Golang (linguagem usada para a criação deste sistema) instalado em sua maquina
- Possuir uma conta no MongoDB Atlas
  - MongoDB Atlas é um banco de dados NoSQL disponível online para armazenar e gerenciar dados.

## Instalação do sistema

Para iniciar faça um clone deste repositorio com o seguinte código em seu **Git Bash**:

`git clone https://github.com/RianAlvesAce/gestao_online`

Após o clone ser finalizado, crie um arquivo `.env` na raiz do projeto e digite o seguinte:

```
CONNECTION_STRING="sua connection string"
JWTSECRET="sua senha para a geração de tokens jwt"
```

Em seguida abra o seu terminal dentro da pasta criada e execute o código abaixo:

`go run ./main.go`

E pronto, seu servidor será iniciado na porta :8080 de seu localhost.

## Trabalho em andamento 🚧

No momento estou trabalhando no desenvolvimento do Front-End do projeto, abaixo deixarei o link para o acesso do protótipo em Figma:

![pedro](https://media.tenor.com/ZAMoMuQgf9UAAAAM/mapache-pedro.gif)