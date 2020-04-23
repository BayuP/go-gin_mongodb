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
			c.Set("credUser", cred.ID)
			c.Next()

		}

	}
}
