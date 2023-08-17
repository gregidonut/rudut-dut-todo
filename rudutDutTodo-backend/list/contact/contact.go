package contact

import (
	"context"
	"fmt"
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/todo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func PingMongo(uri MongoURI) (*mongo.Client, error) {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(uri.ToString()),
	)
	if err != nil {
		return &mongo.Client{},
			fmt.Errorf("%v: %v\n", MongoPingErr.Error(), err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return &mongo.Client{},
			fmt.Errorf("%v: %v\n", MongoPingErr.Error(), err)
	}

	return client, nil
}

// GetList is responsible for the READ method of the api
func GetList(DBCont DBContainer) ([]todo.Todo, error) {
	mongoURIenvVal, err := GetMongoUriFromEnv()
	if err != nil {
		return []todo.Todo{},
			fmt.Errorf("%v: %v", MongoContactErr, err)
	}
	client, err := PingMongo(*mongoURIenvVal)
	if err != nil {
		return []todo.Todo{},
			fmt.Errorf("%v: %v", MongoContactErr, err)
	}

	todoLists := client.Database(DBCont.DBName).Collection(DBCont.CollectionName)

	// retrieve all the documents that match the filter
	cursor, err := todoLists.Find(context.TODO(), bson.D{})
	if err != nil {
		return []todo.Todo{},
			fmt.Errorf("%v: %v", MongoContactErr, err)
	}

	// convert the cursor result to bson
	var bsonResults []bson.M
	err = cursor.All(context.TODO(), &bsonResults)
	if err != nil {
		return []todo.Todo{},
			fmt.Errorf("%v: %v", MongoContactErr, err)
	}

	// convert []bson.M to []todo.Todo
	var todoResults []todo.Todo
	for _, bsonM := range bsonResults {
		todoRes, err := BsonMToTodo(bsonM)
		if err != nil {
			return []todo.Todo{},
				fmt.Errorf("%v: %v", MongoContactErr, err)
		}
		todoResults = append(todoResults, todoRes)
	}

	return todoResults, nil
}
