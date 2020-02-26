package router

import (
	"github.com/Planck1858/yesya-coin/database"
	"github.com/Planck1858/yesya-coin/models"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"net/http"
)

func GetIndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Yesya-coin is running!")
	}
}

func GetBlocksHandler(bc *models.Blockchain, db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		value := bc.GetBlocks()
		spew.Dump(value)

		res := database.SelectAllBlocks(db)
		c.JSON(200, res)
	}
}

func PostMineBlockHandler(bc *models.Blockchain, db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := c.Request.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			panic(err)
		}

		bc.NewBlock(string(data))
		database.NewBlock(db, bc)

		c.JSON(200, gin.H{
			"body": string(data),
		})
	}
}
