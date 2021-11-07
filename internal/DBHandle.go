package internal

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const dbFile = "pass.db"

func DBInit() bool {
	if FileExists() == false {
		if CreateFile() == true {
			return true
		} else {
			return false
		}
	} else {
		return true
	}
}

func FileExists() bool {
	if _, err := os.Stat(dbFile); err == nil {
		return true //TODO:try querying
	} else if errors.Is(err, os.ErrNotExist) {
		return CreateFile()
	} else {
		log.Fatal("Failure in database initialization check.")
		return false
	}
}

func CreateFile() bool {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE PASSWORDS( 
	USER TEXT NOT NULL UNIQUE, 
	PASS TEXT NOT NULL, 
	DT_CREATED TEXT, 
	DT_UPDATED TEXT, 
	PRIMARY KEY(USER));
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return false
	}

	return true
}
