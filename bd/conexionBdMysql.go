package bd

import (
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

/*MysqlConex es el objeto de conexión a la BD */
var MysqlConex = conectarBd()

func conectarBd() (db *sql.DB) {
	//dbDriver := "mysql"
	//dbUser := "root"
	//dbPass := "ste24598ab"
	//dbName := "guardia"
	//db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	db, err := sql.Open("mysql", "root:ste24598ab@tcp(192.168.1.4:3306)/guardia")
	if err != nil {
		panic(err.Error())
	}
	log.Println("Conexión Exitosa con la BD")
	return db
}
