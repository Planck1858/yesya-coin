package router

import (
	"github.com/Planck1858/yesya-coin/models"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	//"github.com/jackc/pgx"
	"io/ioutil"
	"net/http"
)

func GetIndexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Yesya-coin is running!")
}

func GetBlocksHandler(blockchain *models.Blockchain) gin.HandlerFunc {
	return func(c *gin.Context) {
		value := blockchain.GetBlockchain()
		c.BindJSON(&value)
		c.JSON(200, value)

		spew.Dump(blockchain)
	}
}

func PostMineBlockHandler(blockchain *models.Blockchain) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := c.Request.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			panic(err)
		}
		blockchain.NewBlock(string(data))

		c.JSON(200, gin.H{
			"body": string(data),
		})
	}
}
