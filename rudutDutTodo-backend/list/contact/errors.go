package contact

import "errors"

var (
	MongoEnvVarNotDeclaredErr = errors.New("MONGO_URI environment variable is not declared")
	MongoPingErr              = errors.New("error pinging mongoDb Instance")
	MissingDBInfoErr          = errors.New("all fields contact.DBContainer must be filled out")
	MongoContactErr           = errors.New("error contacting mongodb")
)
