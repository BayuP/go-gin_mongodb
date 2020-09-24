package v1

import (
	"context"
	helper "go-gin_mongodb/helpers"
	"go-gin_mongodb/resource/models"
	reqModel "go-gin_mongodb/resource/requestModel/v1"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	uuid "github.com/google/uuid"
)

var collectionTransaction *mongo.Collection

//TransactionsCollections ...
func TransactionsCollections(m *mongo.Database) {
	collectionTransaction = m.Collection("transaction")
}

//CreateTransactions ...
func CreateTransactions(id string, req *reqModel.TransactionReq) map[string]interface{} {

	detailTrans := []models.DetailTransactions{}
	var element models.DetailTransactions
	totalPrice := 0
	for _, v := range req.TransactionDetailReq {
		element.Base = models.Base{
			CreatedTime: time.Now(),
			CreatedBy:   id,
		}
		element.ID = uuid.New().String()
		element.ProductID = v.ProductID
		element.Price = v.Price
		totalPrice += v.Price
		detailTrans = append(detailTrans, element)
	}
	newTransaction := models.Transactions{
		ID:         uuid.New().String(),
		CustomerID: req.CustomerID,
		CashierID:  id,
		Base: models.Base{
			CreatedTime: time.Now(),
			CreatedBy:   id,
		},
		DetailTransactions: detailTrans,
		TotalPrice:         totalPrice,
	}

	_, err := collectionTransaction.InsertOne(context.TODO(), newTransaction)

	if err != nil {
		log.Printf("Error when inserting new users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	reponse := helper.Message(http.StatusCreated, "Succesfull create Role")
	reponse["data"] = newTransaction
	return reponse
}

//GetAllRoles ...
// func GetAllRoles() map[string]interface{} {
// 	filter := bson.M{"base.deletedby": ""}

// 	// userdb := us.User
// 	result := []resModel.RoleResponse{}
// 	//products := []*models.Products{}
// 	//get all user from db
// 	cursor, err := collectionRoles.Find(context.TODO(), filter)

// 	if err != nil {
// 		log.Printf("Error when getting all user %v\n", err)
// 		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
// 		response["data"] = nil
// 		return response
// 	}

// 	for cursor.Next(context.TODO()) {
// 		var roles *models.Roles
// 		cursor.Decode(&roles)
// 		// productResponse := resModel.ProductResponse{
// 		// 	Name:  product.Name,
// 		// 	Price: product.Price,
// 		// }
// 		// products = append(products, product)
// 		roleRes := resModel.RoleResponse{
// 			ID:   roles.ID,
// 			Name: roles.Name,
// 		}
// 		result = append(result, roleRes)
// 	}

// 	// for _, v := range products {
// 	// 	prodRes := resModel.ProductResponse{
// 	// 		ID:    v.ID,
// 	// 		Name:  v.Name,
// 	// 		Price: v.Price,
// 	// 	}
// 	// 	result = append(result, prodRes)
// 	// }

// 	reponse := helper.Message(http.StatusOK, "Succesfull get All user")
// 	reponse["data"] = result
// 	return reponse
// }
