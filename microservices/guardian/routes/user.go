package routes

import (
	"fmt"
	"guardian/api/models"
	"guardian/api/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func createUser(context *gin.Context) {
	var user models.User
	//extract user details from context
	err := context.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err.Error())
		context.JSON(400, gin.H{"message": "bad request - could not parse data"})
		return
	}
	// validate user input
	validate := validator.New()
	err = validate.Struct(user)

	if err != nil {
		fmt.Println(err.Error())
		context.JSON(400, gin.H{"message": "bad request - input data validation failed"})
		return
	}
	// register new user
	err = user.Create()

	if err != nil {
		if strings.Contains(err.Error(), "Error 1062 (23000)") && strings.Contains(err.Error(), "user.display_name") {
			context.JSON(409, gin.H{"message": "display name already in use - could not save user"})
			return
		} else if strings.Contains(err.Error(), "Error 1062 (23000)") && strings.Contains(err.Error(), "user.email") {
			context.JSON(409, gin.H{"message": "email already in use - could not save user"})
			return
		} else {
			fmt.Println(err.Error())
			context.JSON(500, gin.H{"message": "internal server error - could not save user"})
			return
		}
	}

	// user created
	context.JSON(201, gin.H{"message": "created"})
}

func loginUser(context *gin.Context) {
	var user models.User
	//extract user details from context
	err := context.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err.Error())
		context.JSON(400, gin.H{"message": "bad request - could not parse data"})
		return
	}
	// validate user input
	validate := validator.New()
	err = validate.Struct(user)

	if err != nil {
		fmt.Println(err.Error())
		context.JSON(400, gin.H{"message": "bad request - input data validation failed"})
		return
	}
	// validate user credentials
	err = user.ValidateCredentials()

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			context.JSON(401, gin.H{"message": "unauthorised - user not found"})
			return
		} else {
			context.JSON(401, gin.H{"message": "unauthorised - input data validation failed"})
			return
		}
	}
	// check to see if the user is enabled - add code back in admin functionality created
	// if user.InUse == 0 {
	// 	context.JSON(401, gin.H{"message": "unauthorised - user disabled"})
	// 	return
	// }
	// generate token
	token, err := utils.GenerateToken(user.ID, user.DisplayName, user.Email, user.Roles)

	if err != nil {
		context.JSON(500, gin.H{"message": " internal server error - Could not generate token."})
		return
	}
	// all good return token
	context.JSON(200, gin.H{"message": "ok", "token": token})
}

func getUser(context *gin.Context) {
	// get user id from the session context
	userId := context.GetInt64("userId")

	if userId == 0 {
		context.JSON(400, gin.H{"message": "bad request - invalid token"})
		return
	}

	// get user details
	user, err := models.GetUserById(userId)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			context.JSON(404, gin.H{"message": "not found - user not found"})
			return
		} else {
			context.JSON(500, gin.H{"message": "internal server error"})
			return
		}
	}
	// all good return user
	context.JSON(200, gin.H{"message": "ok", "user": user})
}

func updateDisplayName(context *gin.Context) {
	// get user id from the session context
	userId := context.GetInt64("userId")

	if userId == 0 {
		context.JSON(400, gin.H{"message": "bad request - invalid token"})
		return
	}
	var displayName models.DisplayName
	// extract display name to the display name struct
	err := context.ShouldBindJSON(&displayName)

	if err != nil {
		fmt.Println(err.Error())
		context.JSON(400, gin.H{"message": "bad request - could not parse data"})
		return
	}
	// validate user input
	validate := validator.New()
	err = validate.Struct(displayName)

	if err != nil {
		fmt.Println(err.Error())
		context.JSON(400, gin.H{"message": "bad request - input data validation failed"})
		return
	}
	// update user display name
	err = displayName.UpdateDisplayNameById(userId)

	if err != nil {
		fmt.Println(err.Error())
		context.JSON(500, gin.H{"message": "internal server error - could not save display name"})
		return
	}
	// all good return display name
	context.JSON(200, gin.H{"message": "ok", "displayName": displayName.DisplayName})

}

func updateEmail(context *gin.Context) {
	// get user id from the session context
	userId := context.GetInt64(("userId"))

	if userId == 0 {
		context.JSON(400, gin.H{"message": "bad request - invalid token"})
		return
	}
	var email models.Email
	// extract email to the Email struct
	err := context.ShouldBindJSON(&email)

	if err != nil {
		fmt.Println(err.Error())
		context.JSON(400, gin.H{"message": "bad request - could not parse data"})
		return
	}
	// validate user user input
	validate := validator.New()
	err = validate.Struct(email)

	if err != nil {
		fmt.Println(err.Error())
		context.JSON(400, gin.H{"message": "bad request - input data validation failed"})
		return
	}
	// update user email address
	err = email.UpdateEmailById(userId)

	if err != nil {
		if err.Error() == "UNIQUE constraint failed: user.email" {
			context.JSON(409, gin.H{"message": "conflict - could not save user"})
			return
		} else {
			fmt.Println(err.Error())
			context.JSON(500, gin.H{"message": "internal server error - could not save user"})
			return
		}
	}
	// all good return email address
	context.JSON(200, gin.H{"message": "ok", "email": email.Email})
}

func updatePassword(context *gin.Context) {
	// get user id from the session context
	userId := context.GetInt64(("userId"))

	if userId == 0 {
		context.JSON(400, gin.H{"message": "bad request - invalid token"})
		return
	}
	var password models.Password
	// extract email to the Email struct
	err := context.ShouldBindJSON(&password)

	if err != nil {
		fmt.Println(err.Error())
		context.JSON(400, gin.H{"message": "bad request - could not parse data"})
		return
	}
	// validate user user input
	validate := validator.New()
	err = validate.Struct(password)

	if err != nil {
		fmt.Println(err.Error())
		context.JSON(400, gin.H{"message": "bad request - input data validation failed"})
		return
	}
	// update user email address
	err = password.UpdatePasswordById(userId)

	if err != nil {

		fmt.Println(err.Error())
		context.JSON(500, gin.H{"message": "internal server error - could not save user"})
		return
	}
	// all good return 200
	context.JSON(200, gin.H{"message": "ok"})
}
