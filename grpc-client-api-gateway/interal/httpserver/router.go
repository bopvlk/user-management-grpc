package httpserver

import (
	"github.com/gin-gonic/gin"
)

func initRouter(server *httpServer) *gin.Engine {
	router := gin.Default()

	router.POST("/sign-up", server.SignUp)
	router.POST("/sign-in", server.SignIn)

	restrictedGroup := router.Group("/restricted")
	restrictedGroup.GET("/user/:id", server.GetOneUser)

	return router
}
