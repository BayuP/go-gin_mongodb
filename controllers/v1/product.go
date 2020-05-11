package v1

import (
	helpers "go-gin_mongodb/helpers"
	req "go-gin_mongodb/resource/requestModel/v1"
	v1s "go-gin_mongodb/services/v1"

	"github.com/gin-gonic/gin"
)

//CreateProduct ....
func CreateProduct(c *gin.Context) {
	idUser := c.MustGet("credUser").(string)
	var req *req.CreateProdReq
	c.BindJSON(&req)

	response := v1s.CreateProduct(idUser, req)

	helpers.Respond(c.Writer, response)
}

//GetAllProduct ...
func GetAllProduct(c *gin.Context) {

	response := v1s.GetAllProduct()

	helpers.Respond(c.Writer, response)
}

//GetProduct ..
func GetProduct(c *gin.Context) {
	productID := c.Query("id")
	response := v1s.GetProductByID(productID)

	helpers.Respond(c.Writer, response)
}

//UpdateProductByID ...
func UpdateProductByID(c *gin.Context) {
	idUser := c.MustGet("credUser").(string)
	var req *req.UpdateProdReq
	c.BindJSON(&req)
	response := v1s.UpdateProduct(idUser, req)

	helpers.Respond(c.Writer, response)
}

//SoftDeleteByID ...
func SoftDeleteByID(c *gin.Context) {
	idUser := c.MustGet("credUser").(string)
	productID := c.Query("id")
	response := v1s.DeleteProductByID(idUser, productID)

	helpers.Respond(c.Writer, response)
}
