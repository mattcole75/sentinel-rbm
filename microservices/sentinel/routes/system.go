package routes

import (
	"fmt"
	"sentinel/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// create new system
func createSystem(context *gin.Context) {
	var system models.System
	// extract system details from context
	err := context.ShouldBindJSON(&system)
	if err != nil {
		context.JSON(400, gin.H{"message": "bad request - could not parse data"})
		return
	}
	// validate user input
	validate := validator.New()
	err = validate.Struct(system)
	if err != nil {
		context.JSON(400, gin.H{"message": "bad request - input data validation failure"})
		return
	}
	// save new system
	err = system.Create()
	if err != nil {
		context.JSON(500, gin.H{"message": "internal server error - could not save system"})
		return
	}
	// success systems saved
	context.JSON(201, gin.H{"message": "created"})
}

// get systems
func getSystems(context *gin.Context) {
	// fetch all systems
	systems, err := models.GetSystems()
	if err != nil {
		context.JSON(500, gin.H{"message": "internal server error - could not fetch systems"})
		fmt.Println(err.Error())
		return
	}
	// ok return systems
	context.JSON(200, systems)
}

// get system by id
func getSystem(context *gin.Context) {
	// extract the system id from the context param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": "bad request - invalid reference"})
		return
	}
	// fetch the system
	system, err := models.GetSystemById(id)
	if err != nil {
		context.JSON(500, gin.H{"message": "internal server error - could not fetch system"})
		return
	}
	// ok return system
	context.JSON(200, system)
}

// update system
func updateSystem(context *gin.Context) {
	// extract the system id from the context param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": "bad request - invalid reference"})
		return
	}

	// update the system model
	var updatedSystem models.System
	err = context.ShouldBindJSON(&updatedSystem)
	if err != nil {
		fmt.Println(err.Error())
		context.JSON(400, gin.H{"message": "bad request - could not parse data"})
		return
	}
	// update id
	updatedSystem.ID = id
	// update system
	err = updatedSystem.Update()
	if err != nil {
		context.JSON(500, gin.H{"message": "internal server error - could not update system"})
		return
	}
	// ok
	context.JSON(200, updatedSystem)

}

// delete system
func deleteSystem(context *gin.Context) {
	// extract the system id from the context param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": "bad request - invalid reference"})
		return
	}
	// get system from database
	system, err := models.GetSystemById(id)
	if err != nil {
		context.JSON(404, gin.H{"message": "not found - could not fetch system"})
		return
	}
	// delete event
	err = system.Delete()
	if err != nil {
		context.JSON(500, gin.H{"message": "internal server error - could not delete requirement"})
		return
	}
	// ok
	context.JSON(200, system)
}
