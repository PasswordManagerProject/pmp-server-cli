package main

import (
	"database/sql"
	"log"
	"pmp-server/internal/DBHandle"
)

func main() {
	var db *sql.DB

	if DBHandle.DBInit(&db) == false {
		log.Fatal("Failure in DB initializaion.")
		return
	}

	DBHandle.CloseDB(&db)
}
