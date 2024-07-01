package mongodb

import (
	"context"
	"cqrs_go/domain/generic"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoEventRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (repo *MongoEventRepository) Save(event generic.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := repo.collection.InsertOne(ctx, ToMongoEvent(&event))
	return err
}

func (repo *MongoEventRepository) FindByCorrelationID(correlationID string) (*[]generic.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"correlation_id": correlationID}
	cursor, err := repo.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var events []generic.Event
	for cursor.Next(ctx) {
		var mongoEvent MongoEvent
		if err := cursor.Decode(&mongoEvent); err != nil {
			return nil, err
		}
		events = append(events, ToDomainEvent(mongoEvent))
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &events, nil
}

func (repo MongoEventRepository) FindByCategory(category string) (*[]generic.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"category": category}
	cursor, err := repo.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var events []generic.Event
	for cursor.Next(ctx) {
		var event generic.Event
		if err := cursor.Decode(&event); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &events, nil
}
