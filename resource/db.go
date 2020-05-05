package resource

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	service "go-gin_mongodb/services/v1"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//Connect func to db
func Connect() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	clientOptions := options.Client().ApplyURI("mongodb://" + dbHost + ":" + dbPort)
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		print(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("cant connect to db", err)
	} else {
		log.Println("Connected")
	}

	db := client.Database(dbName)
	service.UserCollections(db)
	service.ProductsCollections(db)
	return
}
