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

	user := r.Group("/api/v1/user")
	product := r.Group("/api/v1/product")

	user.POST("/login", v1Controller.LoginUser)

	user.Use(middleware.AuthMiddlewares())
	{

		user.GET("/all", v1Controller.GetAllUsers)
		//v1.GET("/user", v1Controller.GetUser)
		user.POST("/", v1Controller.CreateUser)
		user.PUT("/", v1Controller.UpdateUser)
		user.DELETE("/", v1Controller.DeleteUser)
		user.GET("/", v1Controller.GetUser)
		user.POST("/changepass", v1Controller.ChangePassword)
	}

	product.Use(middleware.AuthMiddlewares())
	{
		product.POST("/create_product", v1Controller.CreateProduct)
		product.GET("/all", v1Controller.GetAllProduct)
		product.GET("/", v1Controller.GetProduct)
		product.PUT("/", v1Controller.UpdateProductByID)
		product.DELETE("/", v1Controller.SoftDeleteByID)
	}

	return r
}
