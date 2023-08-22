package crud

import (
	"context"
	"fmt"
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/contact"
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/todo"
	"go.mongodb.org/mongo-driver/bson"
)

// GetList is responsible for the READ method of the api
func GetList(DBCont contact.DBContainer) ([]todo.Todo, error) {
	mongoURIenvVal, err := contact.GetMongoUriFromEnv()
	if err != nil {
		return []todo.Todo{},
			fmt.Errorf("%v: %v", GetErr, err)
	}
	client, err := contact.PingMongo(*mongoURIenvVal)
	if err != nil {
		return []todo.Todo{},
			fmt.Errorf("%v: %v", GetErr, err)
	}

	todoLists := client.Database(DBCont.DBName).Collection(DBCont.CollectionName)

	// retrieve all the documents that match the filter
	cursor, err := todoLists.Find(context.TODO(), bson.D{})
	if err != nil {
		return []todo.Todo{},
			fmt.Errorf("%v: %v", GetErr, err)
	}

	// convert the cursor result to bson
	var bsonResults []bson.M
	err = cursor.All(context.TODO(), &bsonResults)
	if err != nil {
		return []todo.Todo{},
			fmt.Errorf("%v: %v", GetErr, err)
	}

	// convert []bson.M to []todo.Todo
	var todoResults []todo.Todo
	for _, bsonM := range bsonResults {
		todoRes, err := BsonMToTodo(bsonM)
		if err != nil {
			return []todo.Todo{},
				fmt.Errorf("%v: %v", GetErr, err)
		}
		todoResults = append(todoResults, todoRes)
	}

	return todoResults, nil
}

// AddTodo is responsible for the Create method of the api
func AddTodo(todoObj todo.Todo) error {

	return nil
}
