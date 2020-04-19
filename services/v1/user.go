package v1

import (
	"context"
	helper "go-gin_mongodb/helpers"
	"go-gin_mongodb/resource/models"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

//UserService get struct from models
type UserService struct {
	User models.User
}

//UserCollections to get user collection
func UserCollections(m *mongo.Database) {
	collection = m.Collection("user")
}

//GetAll ...
func (us *UserService) GetAll() map[string]interface{} {
	// userdb := us.User
	usersresponse := []*models.User{}
	//get all user from db
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error when getting all user %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = usersresponse
		return response
	}

	for cursor.Next(context.TODO()) {
		var user *models.User
		cursor.Decode(&user)
		usersresponse = append(usersresponse, user)
	}

	reponse := helper.Message(http.StatusOK, "Succesfull get All user")
	reponse["data"] = usersresponse
	return reponse
}
