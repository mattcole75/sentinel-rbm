package middleware

import (
	"encoding/json"
	"guardian/api/utils"
	"slices"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	// extract the authorisation token from the request header
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		// end transaction with abort
		context.AbortWithStatusJSON(401, gin.H{"message": "not authorised"})
		return
	}
	// verify given token and get user id
	userId, roles, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(401, gin.H{"message": "not authorised"})
		return
	}
	// add the user id to the session context
	context.Set("userId", userId)
	context.Set("roles", roles)

	// execute any pending handlers
	context.Next()
}

func Authorise(context *gin.Context) {

	// pull roles together starting with required role coming in from the requesting api/microservice then the roles the user has been assigned
	var userRoles []string
	requiredRole := context.Request.Header.Get("role")
	_ = json.Unmarshal([]byte(context.GetString("roles")), &userRoles)

	// check to see if the user has the required role
	if slices.Contains(userRoles, requiredRole) {
		context.Next()
	} else {
		context.AbortWithStatusJSON(401, gin.H{"message": "not authorised"})
		return
	}
}
