# ws-go
## Exemplo de webservice escrito em Golang para retorno JSON de informações de logradouros baseado no CEP informado.

### Premissas

Necessário ter o Golang instalado bem como o SQLite3 e o driver SQLite para o Go.

Para instalar o driver SQLite utilizado, comande:

```
$ go get -u github.com/mattn/go-sqlite3
```

O banco de dados é criado e populado durante a execução do programa, no entanto, apenas
alguns CEPs (reais) são inseridos. Segue tabela abaixo com CEPs válidos para os testes:

|   CEP    | Logradouro                   | Localidade/UF            |
|----------|------------------------------|--------------------------|
| 05426200 | Av. Brigadeiro Faria Lima    | São Paulo/SP             |
| 01311000 | Av. Paulista                 | São Paulo/SP             |
| 01415000 | Rua Bela Cintra              | São Paulo/SP             |
| 13010211 | Av. Orozimbo Maia            | Campinas/SP              |
| 13025085 | Rua Barreto Leme             | Campinas/SP              |
| 14010080 | Rua Álvares Cabral           | Ribeirão Preto/SP        |
| 13400560 | Av. Independência            | Piracicaba/SP            |
| 13465280 | Rua Rui Barbosa              | Americana/SP             |
| 13450031 | Av. Monte Castelo            | Santa Bárbara d'Oeste/SP |
| 88020185 | Rua Dr. Jorge Luz Fontes     | Florianópolis/SC         |
| 89201510 | Rua Professora Laura Andrade | Joinville/SC             |
| 90020122 | Rua Dr. Flores               | Porto Alegre/RS          |
| 80010090 | Travessa Frei Caneca         | Curitiba/PR              |

### Exemplo de URL para o uso do webservice

```
http://localhost:8090/CEP/01311000
```

