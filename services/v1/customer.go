package v1

import (
	"context"
	"fmt"
	helper "go-gin_mongodb/helpers"
	"go-gin_mongodb/resource/models"
	requestModel "go-gin_mongodb/resource/requestModel/v1"
	resModel "go-gin_mongodb/resource/responseModel/v1"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

//GetAllCustomer ...
func GetAllCustomer() map[string]interface{} {
	filter := bson.M{"base.deletedby": ""}

	// userdb := us.User
	result := []resModel.CustomerRes{}
	//products := []*models.Products{}
	//get all user from db
	cursor, err := collectionCustomers.Find(context.TODO(), filter)

	if err != nil {
		log.Printf("Error when getting all user %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	for cursor.Next(context.TODO()) {
		var customer *models.Customers
		cursor.Decode(&customer)
		cusRes := resModel.CustomerRes{
			ID:      customer.ID,
			Name:    customer.Name,
			Phone:   customer.Phone,
			Address: customer.Address,
			Email:   customer.Email,
		}
		result = append(result, cusRes)
	}

	reponse := helper.Message(http.StatusOK, "Succesfull get All Customer")
	reponse["data"] = result
	return reponse
}

// GetCustomerByID ...
func GetCustomerByID(id string) map[string]interface{} {
	filter := bson.M{"$and": []bson.M{
		bson.M{"id": id},
		bson.M{"base.deletedby": ""},
	}}

	customer := models.Customers{}
	fmt.Println(filter)
	err := collectionCustomers.FindOne(context.TODO(), filter).Decode(&customer)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			response := helper.Message(http.StatusNotFound, "Not found document")
			response["data"] = nil
			return response
		}
		log.Printf("Error when get Customer : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	customerResponse := resModel.CustomerRes{
		ID:      customer.ID,
		Name:    customer.Name,
		Phone:   customer.Phone,
		Address: customer.Address,
		Email:   customer.Email,
	}

	reponse := helper.Message(http.StatusOK, "Succesfull Get Customer")
	reponse["data"] = customerResponse
	return reponse
}
