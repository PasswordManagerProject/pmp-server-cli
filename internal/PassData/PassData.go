package PassData

import (
	"log"
	"time"
)

type PassData struct {
	init    bool
	ID      string
	User    string
	Pass    string
	DateCtd time.Time
	DateUpd time.Time
}

func IsInit(data PassData) bool {
	return data.init
}

func CreatePassObj(ID string, user string, pass string, dateCtd time.Time, dateUpd time.Time) PassData {
	var data PassData
	var err error

	data.init = false
	data.ID = ID
	data.User = user
	data.Pass = pass

	data.DateCtd = dateCtd
	data.DateUpd = dateUpd

	if err != nil {
		log.Fatal(err)
		data.init = false
		return data
	}

	data.init = true
	return data
}
