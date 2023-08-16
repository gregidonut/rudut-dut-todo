package contact

import (
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
