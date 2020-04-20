package v1

import (
	helpers "go-gin_mongodb/helpers"
	model "go-gin_mongodb/resource/models"
	v1s "go-gin_mongodb/services/v1"

	"github.com/gin-gonic/gin"
)

//GetAllUsers ..
func GetAllUsers(c *gin.Context) {
	var userService v1s.UserService

	response := userService.GetAll()

	helpers.Respond(c.Writer, response)
}

//CreateUser ....
func CreateUser(c *gin.Context) {
	var userService v1s.UserService
	var user *model.User
	c.BindJSON(&user)

	response := userService.Create(user)

	helpers.Respond(c.Writer, response)
}

//UpdateUser ..
func UpdateUser(c *gin.Context) {
	userID := c.Query("id")
	var user *model.User
	c.BindJSON(&user)

	response := v1s.Update(userID, user)

	helpers.Respond(c.Writer, response)
}

//GetUser ..
func GetUser(c *gin.Context) {
	userID := c.Query("id")

	response := v1s.GetByID(userID)

	helpers.Respond(c.Writer, response)
}

//DeleteUser ..
func DeleteUser(c *gin.Context) {
	userID := c.Query("id")

	response := v1s.DeleteByID(userID)

	helpers.Respond(c.Writer, response)
}
