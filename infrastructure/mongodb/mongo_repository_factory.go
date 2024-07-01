package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DATABASE = "testing"
const COLLECTION = "events"

func NewRepository() *MongoEventRepository {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	collection := client.Database(DATABASE).Collection(COLLECTION)

	return &MongoEventRepository{client: client, collection: collection}
}
