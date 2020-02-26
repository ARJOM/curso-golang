package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "arjom:12345678@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin() // iniciando transação
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")
	_, err = stmt.Exec(1, "Thiago") // chave duplicada

	// É preciso tratar o erro para garantir a transação
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
