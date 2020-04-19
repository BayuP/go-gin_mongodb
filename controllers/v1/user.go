package v1

import (
	helpers "go-gin_mongodb/helpers"
	v1s "go-gin_mongodb/services/v1"

	"github.com/gin-gonic/gin"
)

//GetAllUsers ..
func GetAllUsers(c *gin.Context) {
	var userService v1s.UserService

	response := userService.GetAll()

	helpers.Respond(c.Writer, response)
}
