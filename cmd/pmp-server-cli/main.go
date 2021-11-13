package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"pmp-server/internal/DBHandle"
	"pmp-server/internal/RestAPI"
)

func main() {
	var db *sql.DB

	if DBHandle.DBInit(&db) == false {
		log.Fatal("Failure in DB initializaion.")
		return
	}

	router := gin.Default()
	router.POST("/pmp", RestAPI.InsertRec(&db))
	router.PUT("/pmp", RestAPI.UpdateRec(&db))
	router.Run("localhost:8080")

	DBHandle.CloseDB(&db)
}
