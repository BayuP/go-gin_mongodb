package middleware

import (
	"fmt"
	model "go-gin_mongodb/resource/models"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//AuthMiddlewares ...
func AuthMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.Request.Header.Get("Authorization")
		cred := model.Token{}

		e := godotenv.Load()
		if e != nil {
			fmt.Print(e)
		}
		secretKey := os.Getenv("secret_key")

		_, err := jwt.ParseWithClaims(tokenString, &cred, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			fmt.Println("masuk1")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized user",
				"data":    nil,
			})
		} else {
			if (cred.RoleID != "72848e86-112a-4e40-ad60-6ca2f87dadb7") || (cred.RoleID != "216d86e8-90b3-488b-aa30-659a1af9373c") {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"status":  http.StatusUnauthorized,
					"message": "Unauthorized user",
					"data":    nil,
				})
			}
			c.Set("credUser", cred.ID)
			c.Next()

		}

	}
}

//AuthMiddlewaresSA ...
func AuthMiddlewaresSA() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.Request.Header.Get("Authorization")
		cred := model.Token{}

		e := godotenv.Load()
		if e != nil {
			fmt.Print(e)
		}
		secretKey := os.Getenv("secret_key")

		_, err := jwt.ParseWithClaims(tokenString, &cred, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			fmt.Println("masuk1")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized user",
				"data":    nil,
			})
		} else {
			if cred.RoleID != "216d86e8-90b3-488b-aa30-659a1af9373c" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"status":  http.StatusUnauthorized,
					"message": "Unauthorized user",
					"data":    nil,
				})
			}
			c.Set("credUser", cred.ID)
			c.Next()

		}

	}
}

//AuthMiddlewaresCashier ...
func AuthMiddlewaresCashier() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.Request.Header.Get("Authorization")
		cred := model.Token{}

		e := godotenv.Load()
		if e != nil {
			fmt.Print(e)
		}
		secretKey := os.Getenv("secret_key")

		_, err := jwt.ParseWithClaims(tokenString, &cred, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			fmt.Println("masuk1")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized user",
				"data":    nil,
			})
		} else {
			c.Set("credUser", cred.ID)
			c.Next()

		}

	}
}
