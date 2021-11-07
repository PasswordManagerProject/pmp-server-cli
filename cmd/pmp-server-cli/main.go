package main

import (
	"log"
	"pmp-server/internal"
)

func main() {
	if internal.DBInit() == false {
		log.Fatal("Failure in DB initializaion.")
		return
	}

}
