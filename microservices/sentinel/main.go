package main

import (
	"sentinel/api/config"
	"sentinel/api/db"
	"sentinel/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// load configuration
	config.LoadConfigMap()
	// initialise the database connection
	db.InitialiseDB()
	// initialise the web server
	server := gin.Default() // this is a cookie cutter instance
	//register available routes
	routes.RegisterRoutes(server)
	// start server
	server.Run(":6791") // localhost:6791
}
