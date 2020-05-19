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

var collectionCategories *mongo.Collection

//CategoriesCollections ...
func CategoriesCollections(m *mongo.Database) {
	collectionCategories = m.Collection("categories")
}

//CreateCategories ...
func CreateCategories(id string, req *requestModel.CreateCatReq) map[string]interface{} {

	newCategories := models.Categories{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
		Base: models.Base{
			CreatedTime: time.Now(),
			CreatedBy:   id,
		},
	}

	_, err := collectionCategories.InsertOne(context.TODO(), newCategories)

	if err != nil {
		log.Printf("Error when inserting new users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	reponse := helper.Message(http.StatusCreated, "Succesfull create Category")
	reponse["data"] = newCategories
	return reponse
}

//GetAllCategories ...
func GetAllCategories() map[string]interface{} {
	filter := bson.M{"base.deletedby": ""}

	// userdb := us.User
	result := []resModel.CategoriesResponse{}
	//products := []*models.Products{}
	//get all user from db
	cursor, err := collectionCategories.Find(context.TODO(), filter)

	if err != nil {
		log.Printf("Error when getting all user %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	for cursor.Next(context.TODO()) {
		var category *models.Categories
		cursor.Decode(&category)
		// productResponse := resModel.ProductResponse{
		// 	Name:  product.Name,
		// 	Price: product.Price,
		// }
		// products = append(products, product)
		catRes := resModel.CategoriesResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}
		result = append(result, catRes)
	}

	// for _, v := range products {
	// 	prodRes := resModel.ProductResponse{
	// 		ID:    v.ID,
	// 		Name:  v.Name,
	// 		Price: v.Price,
	// 	}
	// 	result = append(result, prodRes)
	// }

	reponse := helper.Message(http.StatusOK, "Succesfull get All user")
	reponse["data"] = result
	return reponse
}

// GetCategoriesByID ...
func GetCategoriesByID(id string) map[string]interface{} {
	filter := bson.M{"$and": []bson.M{
		bson.M{"id": id},
		bson.M{"base.deletedby": ""},
	}}

	categories := models.Categories{}
	fmt.Println(filter)
	err := collectionCategories.FindOne(context.TODO(), filter).Decode(&categories)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			response := helper.Message(http.StatusNotFound, "Not found document")
			response["data"] = nil
			return response
		}
		log.Printf("Error when get Categories : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	categoriesResponse := resModel.CategoriesResponse{
		ID:          categories.ID,
		Name:        categories.Name,
		Description: categories.Description,
	}

	reponse := helper.Message(http.StatusOK, "Succesfull Get Categories")
	reponse["data"] = categoriesResponse
	return reponse
}

//UpdateCategories ..
func UpdateCategories(id string, category *requestModel.UpdateCatReq) map[string]interface{} {
	//filter := bson.M{""}
	filter := bson.M{"$and": []bson.M{
		bson.M{"id": category.ID},
		bson.M{"base.deletedby": ""},
	}}

	newData := bson.M{
		"$set": bson.M{
			"name":             category.Name,
			"description":      category.Description,
			"base.updatedtime": time.Now(),
			"base.updatedby":   id,
		},
	}
	result, err := collectionCategories.UpdateOne(context.TODO(), filter, newData)

	if err != nil {
		log.Printf("Error when updating product : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	if result.MatchedCount == 0 {
		response := helper.Message(http.StatusNotFound, "Not found Document")
		response["data"] = nil
		return response
	}

	reponse := helper.Message(http.StatusOK, "Succesfull Edit category")
	reponse["data"] = nil
	return reponse
}

//DeleteCategoriesByID ..
func DeleteCategoriesByID(userID string, id string) map[string]interface{} {

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

	result, err := collectionCategories.UpdateOne(context.TODO(), filter, newData)

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

	reponse := helper.Message(http.StatusOK, "Succesfull Delete product")
	reponse["data"] = nil
	return reponse
}
