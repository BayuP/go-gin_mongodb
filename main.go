package main

import (
	"fmt"
	"os"

	"go-gin_mongodb/routers"

	"github.com/joho/godotenv"
)

func main() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	r := routers.SetupRouter()

	port := "8080"

	if len(os.Args) > 1 {
		reqPort := os.Args[1]
		if reqPort != "" {
			port = reqPort
		}
	}

	if port == "" {
		port = "8000"
	}

	type Job interface {
		Run()
	}

	r.Run(":" + port)

}
