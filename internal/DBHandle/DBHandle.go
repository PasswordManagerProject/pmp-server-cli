package DBHandle

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"pmp-server/internal/PassData"
	"time"
)

const dtLayout = "2006-01-02 15:04:05.000"
const configFile = "./configs/config.json"

var dbParams string

type Config struct {
	DBParams string `json:"dbParams"`
}

func DBInit(db **sql.DB) bool {
	var err error

	if !GetConfig() {
		return false
	}

	*db, err = sql.Open("mysql", dbParams)

	if err != nil {
		log.Fatal(err)
		return false
	}

	if !CheckTables(db) {
		return CreateTables(db)
	}

	return true
}

func GetConfig() bool {
	var config Config
	configFile, err := os.Open(configFile)
	defer configFile.Close()
	if err != nil {
		log.Fatal(err)
		return false
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	dbParams = config.DBParams
	if dbParams == "" {
		return false
	}

	return true
}

func CheckTables(db **sql.DB) bool {
	var sqlStmt string
	var num int

	sqlStmt = "SELECT COUNT(*) FROM information_schema.TABLES WHERE (TABLE_NAME = 'PASSWORDS');"

	rows, err := (*db).Query(sqlStmt)
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
		return false
	}

	for rows.Next() {
		err = rows.Scan(&num)
	}

	if num == 1 {
		return true
	} else {
		return false
	}
}

func CreateTables(db **sql.DB) bool {
	var err error

	if err != nil {
		log.Fatal(err)
		return false
	}

	sqlStmt := `
	CREATE TABLE PASSWORDS(
	ID VARCHAR(50) NOT NULL UNIQUE,
	USERNAME VARCHAR(50), 
	PASSWORD VARCHAR(50) NOT NULL, 
	DT_CREATED DATETIME, 
	DT_UPDATED DATETIME, 
	PRIMARY KEY(ID));
	`

	_, err = (*db).Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return false
	}

	return true
}

func CloseDB(db **sql.DB) bool {
	var err error
	err = (*db).Close()

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func Insert(db **sql.DB, data PassData.PassData) bool {
	var sqlStmt string
	var err error

	if !PassData.IsInit(data) {
		log.Fatal("Uninitalized struct provided.")
		return false
	}

	sqlStmt = `
   INSERT INTO PASSWORDS (
   ID, USERNAME, PASSWORD,
   DT_CREATED, DT_UPDATED)
   VALUES ( '%s', '%s', '%s', '%s', '%s')
   `
	sqlStmt = fmt.Sprintf(sqlStmt, data.ID, data.User, data.Pass, data.DateCtd.Format(dtLayout), data.DateUpd.Format(dtLayout))

	_, err = (*db).Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func Update(db **sql.DB, ID string, newPass string) bool {
	var sqlStmt string
	var err error
	var dummy PassData.PassData

	if !Query(db, ID, &dummy) {
		return false
	}

	sqlStmt = `
    UPDATE PASSWORDS
    SET PASSWORD = '%s',
    DT_UPDATED = '%s'
    WHERE ID = '%s';
    `

	sqlStmt = fmt.Sprintf(sqlStmt, newPass, time.Now().Format(dtLayout), ID)

	_, err = (*db).Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func Query(db **sql.DB, ID string, data *PassData.PassData) bool {
	var sqlStmt string
	var retID string
	var username string
	var password string
	var dtCtd time.Time
	var dtUpd time.Time

	sqlStmt = fmt.Sprintf("SELECT * FROM PASSWORDS WHERE ID = '%s'", ID)

	rows, err := (*db).Query(sqlStmt)
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
		return false
	}

	for rows.Next() {
		err = rows.Scan(&retID, &username, &password, &dtCtd, &dtUpd)
		*data = PassData.CreatePassObj(retID, username, password, dtCtd, dtUpd)
		if err != nil {
			log.Fatal(err)
			return false
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func List(db **sql.DB) []string {
	var IDList []string
	var sqlStmt string
	var tempVal string

	sqlStmt = "SELECT ID FROM PASSWORDS"

	rows, err := (*db).Query(sqlStmt)
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
		return IDList
	}

	for rows.Next() {
		err = rows.Scan(&tempVal)
		if err != nil {
			log.Fatal(err)
			return IDList
		}
		IDList = append(IDList, tempVal)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return IDList
	}

	return IDList
}

func Count(db **sql.DB) int {
	var sqlStmt string
	var num int

	sqlStmt = "SELECT COUNT(*) FROM PASSWORDS;"

	rows, err := (*db).Query(sqlStmt)
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
		return -1
	}

	for rows.Next() {
		err = rows.Scan(&num)
	}

	return num
}
