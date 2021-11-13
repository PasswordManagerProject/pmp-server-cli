package RestAPI

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pmp-server/internal/DBHandle"
	"pmp-server/internal/PassData"
	"time"
)

type WebData struct {
	Type    string    `json:"Type"`
	ID      string    `json:"ID"`
	User    string    `json:"UserName"`
	Pass    string    `json:"Password"`
	DateCtd time.Time `json:"DtCreated"`
	DateUpd time.Time `json:"DtUpdated"`
}

func InsertRec(db **sql.DB) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		var webData WebData

		if err := ctx.BindJSON(&webData); err != nil {
			log.Print("Error parsing POST request.")
			return
		}

		data := PassData.CreatePassObj(webData.ID, webData.User, webData.Pass, time.Now(), time.Now())
		if DBHandle.Insert(db, data) {
			ctx.IndentedJSON(http.StatusCreated, webData)
		} else {
			ctx.IndentedJSON(http.StatusBadRequest, webData)
		}
	}
	return gin.HandlerFunc(fn)
}

func UpdateRec(db **sql.DB) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		var webData WebData

		if err := ctx.BindJSON(&webData); err != nil {
			log.Print("Error parsing PUT request.")
			return
		}

		if DBHandle.Update(db, webData.ID, webData.Pass) {
			ctx.IndentedJSON(http.StatusAccepted, webData)
		} else {
			ctx.IndentedJSON(http.StatusBadRequest, webData)
		}
	}
	return gin.HandlerFunc(fn)
}
