package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "690718"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	checkErr(err)

	fmt.Println("Successfully connected!")

	//insertion
	stmt, err := db.Prepare("INSERT INTO people(name, birthday, occupation) VALUES($1,$2,$3) RETURNING id")
	checkErr(err)

	res, err := stmt.Exec("Liu","1993","S")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
