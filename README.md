# go-lang-full-cycle-dependency-injection
Repositório para armazenar o código do módulo de injeção de dependências do curso Go-Expert da Full Cycle

Para gerar os arquivos de injeção de dependência, execute o comando abaixo:
```bash
wire
# ou
go generate wire
```

## Dependências
- [Wire](https://github.com/google/wire)

## Rodando o projeto
```bash
go run main.go wire_gen.go
```