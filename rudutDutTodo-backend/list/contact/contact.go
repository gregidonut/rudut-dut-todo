package contact

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func PingMongo(uri MongoURI) error {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(uri.ToString()),
	)
	if err != nil {
		return fmt.Errorf("%q: %q\n", MongoPingErr.Error(), err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return fmt.Errorf("%q: %q\n", MongoPingErr.Error(), err)
	}

	return nil
}
