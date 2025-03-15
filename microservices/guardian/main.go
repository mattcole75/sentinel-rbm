package main

import (
	"guardian/api/config"
	"guardian/api/db"
	"guardian/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// load configuration
	config.LoadConfigMap()
	// initialise the database
	db.InitialiseDB()
	// initialise the web server
	server := gin.Default() // this is a cookie cutter instance
	// register available route
	routes.RegisterRoutes(server)
	//start server
	server.Run(":8080") // localhost:8080
}
