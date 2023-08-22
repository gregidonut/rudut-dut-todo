package crud

import (
	"encoding/json"
	"fmt"
	"github.com/gregidonut/rudut-dut-todo/rudutDutTodo-backend/list/todo"
	"go.mongodb.org/mongo-driver/bson"
)

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
