package v1

import (
	helpers "go-gin_mongodb/helpers"
	req "go-gin_mongodb/resource/requestModel/v1"
	v1s "go-gin_mongodb/services/v1"

	"github.com/gin-gonic/gin"
)

//CreateCategory ....
func CreateCategory(c *gin.Context) {
	idUser := c.MustGet("credUser").(string)
	var req *req.CreateCatReq
	c.BindJSON(&req)

	response := v1s.CreateCategories(idUser, req)

	helpers.Respond(c.Writer, response)
}

//GetAllCategory ...
func GetAllCategory(c *gin.Context) {

	response := v1s.GetAllCategories()

	helpers.Respond(c.Writer, response)
}

//GetCategory ..
func GetCategory(c *gin.Context) {
	categoryID := c.Query("id")
	response := v1s.GetCategoriesByID(categoryID)

	helpers.Respond(c.Writer, response)
}

//UpdateCategoryByID ...
func UpdateCategoryByID(c *gin.Context) {
	idUser := c.MustGet("credUser").(string)
	var req *req.UpdateCatReq
	c.BindJSON(&req)
	response := v1s.UpdateCategories(idUser, req)

	helpers.Respond(c.Writer, response)
}

//DeleteCatByID ...
func DeleteCatByID(c *gin.Context) {
	idUser := c.MustGet("credUser").(string)
	catID := c.Query("id")
	response := v1s.DeleteCategoriesByID(idUser, catID)

	helpers.Respond(c.Writer, response)
}
