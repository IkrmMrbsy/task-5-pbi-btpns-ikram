
package router

import (
	"github.com/gin-gonic/gin"
	"PBIP_BTPN_SYARIAH/controllers"
	"PBIP_BTPN_SYARIAH/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/users/register", controllers.CreateUser)
	r.POST("/users/login", controllers.Login)
	r.PUT("/users/:userId", middlewares.AuthMiddleware(), controllers.UpdateUser)
	r.DELETE("/users/:userId", middlewares.AuthMiddleware(), controllers.DeleteUser)

	r.POST("/photos", middlewares.AuthMiddleware(), controllers.CreatePhoto)
	r.GET("/photos", controllers.GetPhotos)
	r.PUT("/photos/:photoId", middlewares.AuthMiddleware(), controllers.UpdatePhoto)
	r.DELETE("/photos/:photoId", middlewares.AuthMiddleware(), controllers.DeletePhoto)

	return r
}
