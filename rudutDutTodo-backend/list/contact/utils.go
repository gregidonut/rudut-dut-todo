package contact

import (
	"encoding/json"
	"fmt"
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/todo"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)

const (
	MONGO_URI_ENV_VAR = "MONGO_URI"
)

type MongoURI string

func (m MongoURI) ToString() string {
	return string(m)
}

func GetMongoUriFromEnv() (*MongoURI, error) {
	envValue := os.Getenv(MONGO_URI_ENV_VAR)
	mongoUri := MongoURI(envValue)

	if envValue == "" {
		return &mongoUri, MongoEnvVarNotDeclaredErr
	}

	return &mongoUri, nil
}

type DBContainer struct {
	DBName         string
	CollectionName string
}

func NewDBContainer(DBName, collectionName string) (*DBContainer, error) {
	if DBName == "" || collectionName == "" {
		return &DBContainer{}, MissingDBInfoErr
	}

	return &DBContainer{
		DBName:         DBName,
		CollectionName: collectionName,
	}, nil
}

func BsonMToTodo(bsonM bson.M) (todo.Todo, error) {
	object, err := json.MarshalIndent(bsonM, "", "    ")
	if err != nil {
		return todo.Todo{}, fmt.Errorf("%v: %v\n", BsonMToTodoErr, err)
	}

	rawResultJson := json.RawMessage(object)
	todoResult, err := todo.NewTodo(rawResultJson)
	if err != nil {
		return todo.Todo{}, fmt.Errorf("%v: %v\n", BsonMToTodoErr, err)
	}

	return *todoResult, nil
}
