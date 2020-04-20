package routers

import (
	v1Controller "go-gin_mongodb/controllers/v1"
	db "go-gin_mongodb/resource"

	"github.com/gin-gonic/gin"
)

//SetupRouter ...
func SetupRouter() *gin.Engine {

	r := gin.Default()

	//connecting to db
	db.Connect()
	// Routing endpoint

	v1 := r.Group("/api/v1")

	v1.GET("/users", v1Controller.GetAllUsers)
	v1.GET("/user", v1Controller.GetUser)
	v1.POST("/user", v1Controller.CreateUser)
	v1.PUT("/user", v1Controller.UpdateUser)
	v1.DELETE("/user", v1Controller.DeleteUser)

	return r
}
