package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

//CEP Estrutura para retorno do JSON
type CEP struct {
	ID         int    `json:"id"`
	Logradouro string `json:"logradouro"`
	CEP        string `json:"cep"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
}

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func criaDB(db *sql.DB) {
	exec(db, "drop table if exists tblCEPUF")
	exec(db, `create table tblCEPUF (
			id integer primary key,
			UF text not null unique,
			descricao text not null
	)`)

	exec(db, "drop table if exists tblCEPLocalidade")
	exec(db, `create table tblCEPLocalidade (
		id integer primary key,
		localidade text not null,
		id_UF int,
		foreign key (id_UF)
			references tblCEPUF(id)
				on delete no action
				on update no action
	)`)

	exec(db, "drop table if exists tblcepLog")
	exec(db, `create table tblcepLog (
		id integer primary key,
		logradouro text not null,
		CEP text not null,
		id_localidade int,
		foreign key (id_localidade)
			references tblCEPLocalidade(id)
				on delete no action
				on update no action
	)`)

	stmt, _ := db.Prepare("insert into tblCEPUF(id, UF, descricao) values (?, ?, ?)")
	stmt.Exec(1, "AC", "Acre")
	stmt.Exec(2, "AL", "Alagoas")
	stmt.Exec(3, "AM", "Amazonas")
	stmt.Exec(4, "AP", "Amapá")
	stmt.Exec(5, "BA", "Bahia")
	stmt.Exec(6, "CE", "Ceará")
	stmt.Exec(7, "DF", "Distrito Federal")
	stmt.Exec(8, "ES", "Espírito Santo")
	stmt.Exec(9, "GO", "Goias")
	stmt.Exec(10, "MA", "Maranhão")
	stmt.Exec(11, "MG", "Minas Gerais")
	stmt.Exec(12, "MS", "Mato Grosso do Sul")
	stmt.Exec(13, "MT", "Mato Grosso")
	stmt.Exec(14, "PA", "Pará")
	stmt.Exec(15, "PB", "Paraíba")
	stmt.Exec(16, "PE", "Pernambuco")
	stmt.Exec(17, "PI", "Piauí")
	stmt.Exec(18, "PR", "Paraná")
	stmt.Exec(19, "RJ", "Rio de Janeiro")
	stmt.Exec(20, "RN", "Rio Grande do Norte")
	stmt.Exec(21, "RO", "Rondônia")
	stmt.Exec(22, "RR", "Roraima")
	stmt.Exec(23, "RS", "Rio Grande do Sul")
	stmt.Exec(24, "SC", "Santa Catarina")
	stmt.Exec(25, "SE", "Sergipe")
	stmt.Exec(26, "SP", "São Paulo")
	stmt.Exec(27, "TO", "Tocantins")

	stmt2, _ := db.Prepare("insert into tblCEPLocalidade(id, localidade, id_UF) values (?, ?, ?)")
	stmt2.Exec(1, "São Paulo", 26)
	stmt2.Exec(2, "Campinas", 26)
	stmt2.Exec(3, "Ribeirão Preto", 26)
	stmt2.Exec(4, "Piracicaba", 26)
	stmt2.Exec(5, "Americana", 26)
	stmt2.Exec(6, "Santa Bárbara d'Oeste", 26)
	stmt2.Exec(7, "Florianópolis", 24)
	stmt2.Exec(8, "Joinville", 24)
	stmt2.Exec(9, "Porto Alegre", 23)
	stmt2.Exec(10, "Curitiba", 18)

	stmt3, _ := db.Prepare("insert into tblcepLog(id, logradouro, CEP, id_localidade) values (?, ?, ?, ?)")
	stmt3.Exec(1, "Av. Brigadeiro Faria Lima", "05426200", 1)
	stmt3.Exec(2, "Av. Paulista", "01311000", 1)
	stmt3.Exec(3, "Rua Bela Cintra", "01415000", 1)
	stmt3.Exec(4, "Av. Orozimbo Maia", "13010211", 2)
	stmt3.Exec(5, "Rua Barreto Leme", "13025085", 2)
	stmt3.Exec(6, "Rua Álvares Cabral", "14010080", 3)
	stmt3.Exec(7, "Av. Independência", "13400560", 4)
	stmt3.Exec(8, "Rua Rui Barbosa", "13465280", 5)
	stmt3.Exec(9, "Av. Monte Castelo", "13450031", 6)
	stmt3.Exec(10, "Rua Dr. Jorge Luz Fontes", "88020185", 7)
	stmt3.Exec(11, "Rua Professora Laura Andrade", "89201510", 8)
	stmt3.Exec(12, "Rua Dr. Flores", "90020122", 9)
	stmt3.Exec(13, "Travessa Frei Caneca", "80010090", 10)
}

//CEPHandler analisa o request e delega para a funcao adequada
func CEPHandler(w http.ResponseWriter, r *http.Request) {
	sCEP := strings.TrimPrefix(r.URL.Path, "/CEP/")

	switch {
	case r.Method == "GET" && sCEP != "":
		buscaCEP(w, r, sCEP)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Necessário informar um CEP.... ")
	}
}

func buscaCEP(w http.ResponseWriter, r *http.Request, cep string) {
	db, err := sql.Open("sqlite3", "./dbCEP.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var cepLog CEP

	db.QueryRow("select tcl.id, tcl.logradouro, tco.localidade, tcu.UF, tcl.CEP from tblcepLog tcl join tblCEPLocalidade tco on tco.id = tcl.id_localidade join tblCEPUF tcu on tco.id_UF = tcu.id where tcl.cep = ?", cep).Scan(&cepLog.ID, &cepLog.Logradouro, &cepLog.Localidade, &cepLog.UF, &cepLog.CEP)

	json, _ := json.Marshal(cepLog)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func main() {
	db, err := sql.Open("sqlite3", "./dbCEP.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	criaDB(db)

	http.HandleFunc("/CEP/", CEPHandler)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
