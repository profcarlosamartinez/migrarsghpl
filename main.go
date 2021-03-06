package main

import (
	"database/sql"
	"fmt"

	"github.com/profcarlosamartinez/migrarsghpl/bd"

	_ "github.com/go-sql-driver/mysql"
)

type Persona struct {
	id      int
	nombres string
}

func main() {

	err := bd.MysqlConex.Ping()

	bd.ChequeoConnection()

	if err != nil {
		fmt.Println("Conexión exitosa")
	}

	fmt.Println("Migrar en Go con MySQL")

	db, err := sql.Open("mysql", "root:ste24598ab@tcp(192.168.1.4:3306)/guardia")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, nombres FROM persona")

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var persona Persona
		err := rows.Scan(&persona.id, &persona.nombres)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(persona)
	}

}
