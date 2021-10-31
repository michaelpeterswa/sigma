package mongo

import (
	"context"
	"log"

	"github.com/michaelpeterswa/sigma/backend/cmd/structs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConn struct {
	Client mongo.Client
}

func InitMongo(ctx context.Context, settings structs.Settings) *MongoConn {

	/*
	   Connect to my cluster
	*/
	client, err := mongo.NewClient(options.Client().ApplyURI(settings.MongoURI))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return &MongoConn{
		Client: *client,
	}

}
