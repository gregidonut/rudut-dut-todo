package crud

import "errors"

var (
	BsonMToTodoErr = errors.New("having trouble converting from bson to Todo")
	GetErr         = errors.New("error getting list from mongodb")
)
