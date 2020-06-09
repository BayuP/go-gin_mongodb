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

//GetAllCustomer ...
func GetAllCustomer(c *gin.Context) {

	response := v1s.GetAllCustomer()

	helpers.Respond(c.Writer, response)
}

//GetCustomer ..
func GetCustomer(c *gin.Context) {
	customerID := c.Query("id")
	response := v1s.GetCustomerByID(customerID)

	helpers.Respond(c.Writer, response)
}
