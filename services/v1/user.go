package v1

import (
	"context"
	"fmt"
	helper "go-gin_mongodb/helpers"
	"go-gin_mongodb/resource/models"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	uuid "github.com/google/uuid"
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

//Create ..
func (us *UserService) Create(user *models.User) map[string]interface{} {

	newUser := models.User{
		ID:        uuid.New().String(),
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := collection.InsertOne(context.TODO(), newUser)

	if err != nil {
		log.Printf("Error when inserting new users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	reponse := helper.Message(http.StatusCreated, "Succesfull create user")
	reponse["data"] = newUser
	return reponse
}

//Update ..
func Update(id string, user *models.User) map[string]interface{} {
	newData := bson.M{
		"$set": bson.M{
			"username":   user.Username,
			"password":   user.Password,
			"updated_at": time.Now(),
		},
	}
	fmt.Println(user.Username)
	fmt.Println(id)

	_, err := collection.UpdateOne(context.TODO(), bson.M{"id": id}, newData)

	if err != nil {
		log.Printf("Error when updating users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	reponse := helper.Message(http.StatusOK, "Succesfull Edit user")
	reponse["data"] = newData
	return reponse
}

// GetByID ..
func GetByID(id string) map[string]interface{} {
	user := models.User{}

	err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&user)

	if err != nil {
		log.Printf("Error when get users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	reponse := helper.Message(http.StatusOK, "Succesfull Get user")
	reponse["data"] = user
	return reponse
}

// DeleteByID ..
func DeleteByID(id string) map[string]interface{} {
	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": id})

	if err != nil {
		log.Printf("Error when delete users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	reponse := helper.Message(http.StatusOK, "Succesfull Delete user")
	reponse["data"] = nil
	return reponse
}
