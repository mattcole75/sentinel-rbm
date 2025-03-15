package routes

import (
	"guardian/api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// unprotected routes
	server.POST("/user/create", createUser)
	server.POST("/user/login", loginUser)

	// setup protected routes
	authenticated := server.Group("/")
	// link the following routes to the authenticate middleware
	authenticated.Use(middleware.Authenticate)
	// protected routes
	// authenticated.GET("/user", middleware.Authorise, getUser)
	authenticated.GET("/user", getUser)
	authenticated.PATCH("/user/displayname", updateDisplayName)
	authenticated.PATCH("/user/email", updateEmail)
	authenticated.PATCH("/user/password", updatePassword)
	authenticated.POST("/transaction/authorise", middleware.Authorise)
}
