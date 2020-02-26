package router

import (
	"github.com/Planck1858/yesya-coin/database"
	"github.com/Planck1858/yesya-coin/models"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	bc := models.NewBlockchain()
	db := database.Init("yesya", bc)

	r.GET("/", GetIndexHandler())
	r.GET("/blocks", GetBlocksHandler(bc, db))
	r.POST("/mineBlock", PostMineBlockHandler(bc, db))
	return r
}
