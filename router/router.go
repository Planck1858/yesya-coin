package router

import (
	"github.com/Planck1858/yesya-coin/models"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	bc := models.NewBlockchain()

	r.GET("/", GetIndexHandler)
	r.GET("/blocks", GetBlocksHandler(bc))
	r.POST("/mineBlock", PostMineBlockHandler(bc))
	return r
}
