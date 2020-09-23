package v1

import (
	helpers "go-gin_mongodb/helpers"
	req "go-gin_mongodb/resource/requestModel/v1"
	v1s "go-gin_mongodb/services/v1"

	"github.com/gin-gonic/gin"
)

//GetAllRole ..
func GetAllRole(c *gin.Context) {

	response := v1s.GetAllRoles()

	helpers.Respond(c.Writer, response)
}

//CreateRole ....
func CreateRole(c *gin.Context) {
	// idUser := c.MustGet("credUser").(string)
	idUser := "System"
	var role *req.CreateRoleReq
	c.BindJSON(&role)

	response := v1s.CreateRoles(idUser, role)

	helpers.Respond(c.Writer, response)
}
