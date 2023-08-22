package crud

import "errors"

var (
	BsonMToTodoErr  = errors.New("having trouble converting from bson to Todo")
	MongoContactErr = errors.New("error contacting mongodb")
)
