package routers

import (
	v1Controller "go-gin_mongodb/controllers/v1"
	"go-gin_mongodb/middleware"
	db "go-gin_mongodb/resource"

	"github.com/gin-gonic/gin"
)

//SetupRouter ...
func SetupRouter() *gin.Engine {

	r := gin.Default()

	//connecting to db
	db.Connect()
	// Routing endpoint

	v1 := r.Group("/api/v1/user")

	v1.GET("/all", v1Controller.GetAllUsers)
	//v1.GET("/user", v1Controller.GetUser)
	v1.POST("/", v1Controller.CreateUser)
	v1.PUT("/", v1Controller.UpdateUser)
	v1.DELETE("/", v1Controller.DeleteUser)

	v1.POST("/login", v1Controller.LoginUser)

	v1.Use(middleware.AuthMiddlewares())
	{

		v1.GET("/", v1Controller.GetUser)
		v1.POST("/changepass", v1Controller.ChangePassword)
	}

	return r
}
