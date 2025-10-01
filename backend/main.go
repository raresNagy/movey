package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"

)

type users struct {
	name string
	password string

}



func main() {
	userdb, err:= sql.Open("sqlite", "./users.db");
	if  err != nil {
		log.Fatal(err)
	}

	stmt, err = userdb.Prepare("CREATE")

}
