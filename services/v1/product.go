package v1

import (
	"context"
	helper "go-gin_mongodb/helpers"
	"go-gin_mongodb/resource/models"
	requestModel "go-gin_mongodb/resource/requestModel/v1"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	uuid "github.com/google/uuid"
)

var collectionProduct *mongo.Collection

//ProductsCollections ...
func ProductsCollections(m *mongo.Database) {
	collectionProduct = m.Collection("product")
}

//CreateProduct ...
func CreateProduct(id string, req *requestModel.CreateProdReq) map[string]interface{} {

	newProduct := models.Products{
		ID:    uuid.New().String(),
		Name:  req.Name,
		Price: req.Price,
		Base: models.Base{
			CreatedTime: time.Now(),
			CreatedBy:   id,
		},
	}

	_, err := collectionProduct.InsertOne(context.TODO(), newProduct)

	if err != nil {
		log.Printf("Error when inserting new users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	reponse := helper.Message(http.StatusCreated, "Succesfull create Product")
	reponse["data"] = newProduct
	return reponse
}
