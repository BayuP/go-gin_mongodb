package v1

import (
	"fmt"
	helpers "go-gin_mongodb/helpers"
	model "go-gin_mongodb/resource/models"
	req "go-gin_mongodb/resource/requestModel/v1"
	v1s "go-gin_mongodb/services/v1"

	"github.com/gin-gonic/gin"
)

//GetAllUsers ..
func GetAllUsers(c *gin.Context) {

	response := v1s.GetAll()

	helpers.Respond(c.Writer, response)
}

//CreateUser ....
func CreateUser(c *gin.Context) {
	// idUser := c.MustGet("credUser").(string)
	idUser := "System"
	var user *req.CreateUserReq
	c.BindJSON(&user)

	response := v1s.Create(idUser, user)

	helpers.Respond(c.Writer, response)
}

//UpdateUser ..
func UpdateUser(c *gin.Context) {
	idUser := c.MustGet("credUser").(string)
	var user *req.EditUserReq
	c.BindJSON(&user)

	response := v1s.Update(idUser, user)

	helpers.Respond(c.Writer, response)
}

//GetUser ..
func GetUser(c *gin.Context) {
	//userID := c.Query("id")
	idUser := c.MustGet("credUser").(string)
	response := v1s.GetByID(idUser)
	fmt.Println(idUser)
	helpers.Respond(c.Writer, response)
}

//DeleteUser ..
func DeleteUser(c *gin.Context) {
	userID := c.Query("id")
	idUser := c.MustGet("credUser").(string)
	response := v1s.DeleteByID(idUser, userID)

	helpers.Respond(c.Writer, response)
}

//LoginUser ..
func LoginUser(c *gin.Context) {
	var user *model.User
	c.BindJSON(&user)

	response := v1s.Login(user)

	helpers.Respond(c.Writer, response)
}

//ChangePassword ...
func ChangePassword(c *gin.Context) {
	idUser := c.MustGet("credUser").(string)
	var reqModel *req.ChangePassReqModel
	c.BindJSON(&reqModel)

	response := v1s.ChangePass(idUser, reqModel)

	helpers.Respond(c.Writer, response)

}
