package v1

import (
	"context"
	helper "go-gin_mongodb/helpers"
	"go-gin_mongodb/resource/models"
	reqModel "go-gin_mongodb/resource/requestModel/v1"
	resModel "go-gin_mongodb/resource/responseModel/v1"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	uuid "github.com/google/uuid"
)

var collectionRoles *mongo.Collection

//RolesCollections ...
func RolesCollections(m *mongo.Database) {
	collectionRoles = m.Collection("roles")
}

//CreateRoles ...
func CreateRoles(id string, req *reqModel.CreateRoleReq) map[string]interface{} {

	newRole := models.Roles{
		ID:   uuid.New().String(),
		Name: req.Name,
		Base: models.Base{
			CreatedTime: time.Now(),
			CreatedBy:   id,
		},
	}

	_, err := collectionRoles.InsertOne(context.TODO(), newRole)

	if err != nil {
		log.Printf("Error when inserting new users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	reponse := helper.Message(http.StatusCreated, "Succesfull create Role")
	reponse["data"] = newRole
	return reponse
}

//GetAllRoles ...
func GetAllRoles() map[string]interface{} {
	filter := bson.M{"base.deletedby": ""}

	// userdb := us.User
	result := []resModel.RoleResponse{}
	//products := []*models.Products{}
	//get all user from db
	cursor, err := collectionRoles.Find(context.TODO(), filter)

	if err != nil {
		log.Printf("Error when getting all user %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	for cursor.Next(context.TODO()) {
		var roles *models.Roles
		cursor.Decode(&roles)
		// productResponse := resModel.ProductResponse{
		// 	Name:  product.Name,
		// 	Price: product.Price,
		// }
		// products = append(products, product)
		roleRes := resModel.RoleResponse{
			ID:   roles.ID,
			Name: roles.Name,
		}
		result = append(result, roleRes)
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
