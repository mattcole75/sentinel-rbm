package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/system", createSystem)
	server.GET("/systems", getSystems)
	server.GET("/system/:id", getSystem)
	server.PATCH("/system/:id", updateSystem)
	server.DELETE("/system/:id", deleteSystem)

	server.POST("/requirement", createRequirement)
	server.GET("/requirements", getRequirements)
	server.GET("/requirement/:id", getRequirement)
	server.PATCH("/requirement/:id", updateRequirement)
	server.DELETE("/requirement/:id", deleteRequirement)
}
