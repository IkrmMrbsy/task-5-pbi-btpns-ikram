
package app

import (
	"github.com/gin-gonic/gin"
	"PBIP_BTPN_SYARIAH/database"
)

var (
	router *gin.Engine
)

func Init() {
	router = gin.Default()
	database.InitDB()
}
