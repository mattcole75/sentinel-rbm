package routes

import (
	"sentinel/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func createRequirement(context *gin.Context) {
	var requirement models.Requirement
	// extract requirement details from context
	err := context.ShouldBindJSON(&requirement)
	if err != nil {
		context.JSON(400, gin.H{"message": "bad request - could not parse data"})
		return
	}

	// validate user input
	validate := validator.New()
	err = validate.Struct(requirement)
	if err != nil {
		context.JSON(400, gin.H{"message": "bad request - input data validation failed"})
		return
	}

	// save new requirement
	err = requirement.Create()
	if err != nil {
		context.JSON(500, gin.H{"message": "internal server error - could not save requirement"})
		return
	}

	// success requirement created
	context.JSON(201, gin.H{"message": "created"})
}

func getRequirements(context *gin.Context) {
	// fetch all requirements
	requirements, err := models.GetRequirements()
	if err != nil {
		context.JSON(500, gin.H{"message": "could not fetch requirements"})
		return
	}
	// ok return requirements
	context.JSON(200, requirements)
}

func getRequirement(context *gin.Context) {
	// extract the requirement id from the context params
	requirementId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": "invalid requirement id"})
		return
	}
	// get the requirement by id
	requirement, err := models.GetRequirementById(requirementId)
	if err != nil {
		context.JSON(500, gin.H{"message": "could not fetch requirement"})
	}
	// ok return requirement
	context.JSON(200, requirement)
}

func updateRequirement(context *gin.Context) {
	// extract the requirement id from the context params
	requirementId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": "invalid requirement id"})
		return
	}

	// get requirement from the database
	// requirement, err := models.GetRequirementById(requirementId)
	// if err != nil {
	// 	context.JSON(404, gin.H{"message": "could not fetch requirement"})
	// 	return
	// }

	// update requirement
	var updatedRequirement models.Requirement
	err = context.ShouldBindJSON(&updatedRequirement)
	if err != nil {
		context.JSON(400, gin.H{"message": "could not parse given data"})
		return
	}
	updatedRequirement.ID = requirementId

	// update event
	err = updatedRequirement.Update()
	if err != nil {
		context.JSON(500, gin.H{"message": "could not update requirement"})
		return
	}

	context.JSON(200, updatedRequirement)
}

func deleteRequirement(context *gin.Context) {
	requirementId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": "invalid requirement id"})
		return
	}

	// get requirement from the database
	requirement, err := models.GetRequirementById(requirementId)
	if err != nil {
		context.JSON(404, gin.H{"message": "could not fetch requirement"})
		return
	}

	// delete event
	err = requirement.Delete()
	if err != nil {
		context.JSON(500, gin.H{"message": "could not update requirement"})
		return
	}

	context.JSON(200, requirement)

}
