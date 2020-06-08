package v1

import (
	helpers "go-gin_mongodb/helpers"
	req "go-gin_mongodb/resource/requestModel/v1"
	v1s "go-gin_mongodb/services/v1"

	"github.com/gin-gonic/gin"
)

//CreateCustomer ...
func CreateCustomer(c *gin.Context) {
	idUser := c.MustGet("credUser").(string)
	var req *req.CustomerReq
	c.BindJSON(&req)

	response := v1s.CreateCustomer(idUser, req)

	helpers.Respond(c.Writer, response)
}
