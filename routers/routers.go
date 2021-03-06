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
	category := r.Group("/api/v1/category")
	customer := r.Group("/api/v1/customer")
	role := r.Group("/api/v1/role")
	transaction := r.Group("/api/v1/transaction")

	user.POST("/login", v1Controller.LoginUser)

	user.Use(middleware.AuthMiddlewaresSA())
	{
		user.POST("/", v1Controller.CreateUser)
		user.GET("/all", v1Controller.GetAllUsers)
		role.POST("/", v1Controller.CreateRole)
		role.GET("/all", v1Controller.GetAllRole)
	}

	user.Use(middleware.AuthMiddlewares())
	{

		//v1.GET("/user", v1Controller.GetUser)

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

	category.Use(middleware.AuthMiddlewares())
	{
		category.POST("/create_category", v1Controller.CreateCategory)
		category.GET("/all", v1Controller.GetAllCategory)
		category.GET("/", v1Controller.GetCategory)
		category.PUT("/", v1Controller.UpdateCategoryByID)
		category.DELETE("/", v1Controller.DeleteCatByID)
	}

	customer.Use(middleware.AuthMiddlewaresCashier())
	{
		customer.POST("/create_customer", v1Controller.CreateCustomer)
		customer.GET("/", v1Controller.GetCustomer)
		customer.GET("/all", v1Controller.GetAllCustomer)

	}

	transaction.Use(middleware.AuthMiddlewares())
	{
		transaction.POST("/", v1Controller.CreateTrans)
	}

	return r
}
