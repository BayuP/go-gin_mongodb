package v1

import (
	"context"
	"fmt"
	helper "go-gin_mongodb/helpers"
	"go-gin_mongodb/resource/models"
	requestModel "go-gin_mongodb/resource/requestModel/v1"
	responseModel "go-gin_mongodb/resource/responseModel/v1"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/joho/godotenv"
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
func GetAll() map[string]interface{} {
	filter := bson.M{"base.deletedby": ""}
	// userdb := us.User
	usersresponse := []responseModel.UserResponse{}
	//get all user from db
	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Printf("Error when getting all user %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = usersresponse
		return response
	}

	for cursor.Next(context.TODO()) {
		var user *models.User
		cursor.Decode(&user)
		userRes := responseModel.UserResponse{
			ID:       user.ID,
			Username: user.Username,
		}

		usersresponse = append(usersresponse, userRes)
	}

	reponse := helper.Message(http.StatusOK, "Succesfull get All user")
	reponse["data"] = usersresponse
	return reponse
}

//Create ..
func Create(id string, req *requestModel.CreateUserReq) map[string]interface{} {

	newUser := models.User{
		ID:       uuid.New().String(),
		Username: req.Username,
		Password: req.Password,
		Base: models.Base{
			CreatedTime: time.Now(),
			CreatedBy:   id,
		},
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
func Update(id string, user *requestModel.EditUserReq) map[string]interface{} {

	filter := bson.M{"$and": []bson.M{
		bson.M{"id": user.ID},
		bson.M{"base.deletedby": ""},
	}}

	newData := bson.M{
		"$set": bson.M{
			"username":         user.Username,
			"base.updatedtime": time.Now(),
			"base.updatedby":   id,
		},
	}

	result, err := collection.UpdateOne(context.TODO(), filter, newData)

	if err != nil {
		log.Printf("Error when updating users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	if result.MatchedCount == 0 {
		response := helper.Message(http.StatusNotFound, "Not found Document")
		response["data"] = nil
		return response
	}

	reponse := helper.Message(http.StatusOK, "Succesfull Edit user")
	reponse["data"] = newData
	return reponse
}

// GetByID ..
func GetByID(id string) map[string]interface{} {
	filter := bson.M{"$and": []bson.M{
		bson.M{"id": id},
		bson.M{"base.deletedby": ""},
	}}

	user := models.User{}

	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			response := helper.Message(http.StatusNotFound, "Not found document")
			response["data"] = nil
			return response
		}

		log.Printf("Error when get users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	userResponse := responseModel.UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	reponse := helper.Message(http.StatusOK, "Succesfull Get user")
	reponse["data"] = userResponse
	return reponse
}

// DeleteByID ..
func DeleteByID(userID string, id string) map[string]interface{} {
	filter := bson.M{"$and": []bson.M{
		bson.M{"id": id},
		bson.M{"base.deletedby": ""},
	}}

	newData := bson.M{
		"$set": bson.M{
			"base.deletedtime": time.Now(),
			"base.deletedby":   userID,
		},
	}

	result, err := collection.UpdateOne(context.TODO(), filter, newData)

	if err != nil {
		log.Printf("Error when delete users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	if result.MatchedCount == 0 {
		response := helper.Message(http.StatusNotFound, "Not found Document")
		response["data"] = nil
		return response
	}

	reponse := helper.Message(http.StatusOK, "Succesfull Delete user")
	reponse["data"] = nil
	return reponse
}

//Login ...
func Login(model *models.User) map[string]interface{} {

	filter := bson.M{"$and": []bson.M{
		bson.M{"username": model.Username},
		bson.M{"base.deletedby": ""},
	}}

	filterUser := bson.M{"$and": []bson.M{
		bson.M{"username": model.Username},
		bson.M{"password": model.Password},
	}}

	user := models.User{}

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	secretKey := os.Getenv("secret_key")

	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	fmt.Println(err)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Error get users : %v\n", err)
			response := helper.Message(http.StatusNotFound, "User not found")
			response["data"] = nil
			return response
		}
		log.Printf("Error get users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	errFindUser := collection.FindOne(context.TODO(), filterUser).Decode(&user)
	fmt.Println(err)

	if errFindUser != nil {
		if errFindUser == mongo.ErrNoDocuments {
			log.Printf("Error get users : %v\n", errFindUser)
			response := helper.Message(http.StatusNotFound, "Username & Password not Match")
			response["data"] = nil
			return response
		}
		log.Printf("Error get users : %v\n", errFindUser)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	expiredTime := time.Now().Add(1000 * time.Minute)

	claims := &models.Token{
		Username: user.Username,
		ID:       user.ID,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	fmt.Println(token)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Printf("Error creating jwt users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	response := responseModel.LoginResponse{
		Username: user.Username,
		Token:    tokenString,
		ID:       user.ID,
	}

	reponse := helper.Message(http.StatusOK, "Succesfull Login")
	reponse["data"] = response
	return reponse

}

//ChangePass ...
func ChangePass(id string, req *requestModel.ChangePassReqModel) map[string]interface{} {

	filter := bson.M{"$and": []bson.M{
		bson.M{"id": req.ID},
		bson.M{"password": req.OldPassword},
	}}
	newData := bson.M{
		"$set": bson.M{
			"password":         req.NewPassword,
			"base.updatedtime": time.Now(),
			"base.updatedby":   id,
		},
	}
	fmt.Println(req.NewPassword)
	fmt.Println(req.OldPassword)

	result, err := collection.UpdateOne(context.TODO(), filter, newData)

	if result.MatchedCount == 0 {
		response := helper.Message(http.StatusOK, "Old Password Didnt match")
		response["data"] = nil
		return response
	}

	if err != nil {
		log.Printf("Error when updating users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	response := helper.Message(http.StatusOK, "Successfull Change Password")
	response["data"] = nil
	return response
}
