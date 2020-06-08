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

var collectionCustomers *mongo.Collection

//CustomersCollections ...
func CustomersCollections(m *mongo.Database) {
	collectionCustomers = m.Collection("customer")
}

//CreateCustomer ...
func CreateCustomer(id string, req *requestModel.CustomerReq) map[string]interface{} {

	newCustomer := models.Customers{
		ID:      uuid.New().String(),
		Name:    req.Name,
		Address: req.Address,
		Email:   req.Email,
		Phone:   req.Phone,
		Base: models.Base{
			CreatedTime: time.Now(),
			CreatedBy:   id,
		},
	}

	_, err := collectionCustomers.InsertOne(context.TODO(), newCustomer)

	if err != nil {
		log.Printf("Error when inserting new users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	reponse := helper.Message(http.StatusCreated, "Succesfull create Customer")
	reponse["data"] = newCustomer
	return reponse
}
