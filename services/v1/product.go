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

//GetAllProduct ...
func GetAllProduct() map[string]interface{} {
	filter := bson.M{"base.deletedby": ""}

	// userdb := us.User
	result := []resModel.ProductResponse{}
	//products := []*models.Products{}
	//get all user from db
	cursor, err := collectionProduct.Find(context.TODO(), filter)

	if err != nil {
		log.Printf("Error when getting all user %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	for cursor.Next(context.TODO()) {
		var product *models.Products
		cursor.Decode(&product)
		// productResponse := resModel.ProductResponse{
		// 	Name:  product.Name,
		// 	Price: product.Price,
		// }
		// products = append(products, product)
		prodRes := resModel.ProductResponse{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		}
		result = append(result, prodRes)
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

// GetProductByID ...
func GetProductByID(id string) map[string]interface{} {
	filter := bson.M{"$and": []bson.M{
		bson.M{"id": id},
		bson.M{"base.deletedby": ""},
	}}
	product := models.Products{}
	fmt.Println(filter)
	err := collectionProduct.FindOne(context.TODO(), filter).Decode(&product)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			response := helper.Message(http.StatusNotFound, "Not found document")
			response["data"] = nil
			return response
		}
		log.Printf("Error when get Product : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	productResponse := resModel.ProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}

	reponse := helper.Message(http.StatusOK, "Succesfull Get product")
	reponse["data"] = productResponse
	return reponse
}

//UpdateProduct ..
func UpdateProduct(id string, product *requestModel.UpdateProdReq) map[string]interface{} {
	//filter := bson.M{""}
	filter := bson.M{"$and": []bson.M{
		bson.M{"id": product.ID},
		bson.M{"base.deletedby": ""},
	}}

	newData := bson.M{
		"$set": bson.M{
			"name":             product.Name,
			"price":            product.Price,
			"base.updatedtime": time.Now(),
			"base.updatedby":   id,
		},
	}
	result, err := collectionProduct.UpdateOne(context.TODO(), filter, newData)

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

	reponse := helper.Message(http.StatusOK, "Succesfull Edit product")
	reponse["data"] = nil
	return reponse
}

//DeleteProductByID ..
func DeleteProductByID(userID string, id string) map[string]interface{} {

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

	result, err := collectionProduct.UpdateOne(context.TODO(), filter, newData)

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
