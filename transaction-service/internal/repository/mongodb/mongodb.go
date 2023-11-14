package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type MongoDBImpl struct {
	client mongo.Client
}

func New(ctx context.Context) (MongoDBImpl ,error){

	// todo: move to envs
	uri := "mongodb://admin:password@localhost:27017/?timeoutMS=5000"

	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return MongoDBImpl{}, err
	}

	var result bson.M
	if err := client.Database("admin").RunCommand(ctx, bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return MongoDBImpl{client: *client}, nil
}
