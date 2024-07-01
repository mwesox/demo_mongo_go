package mongodb

import (
	"cqrs_go/domain/generic"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*
Mongo specific representation of the event
*/
type MongoEvent struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	EventType     string             `bson:"event_type"`
	Payload       []byte             `bson:"payload"`
	Timestamp     time.Time          `bson:"timestamp"`
	CorrelationID string             `bson:"correlation_id"`
	Category      string             `bson:"category"`
}

func ToMongoEvent(event *generic.Event) *MongoEvent {
	objectID, err := primitive.ObjectIDFromHex(event.Id)
	if err != nil {
		objectID = primitive.NewObjectID()
	} else {
		panic(err)
	}
	return &MongoEvent{
		ID:            objectID,
		EventType:     event.EventType,
		Payload:       event.Payload,
		Timestamp:     event.Timestamp,
		CorrelationID: event.CorrelationId,
		Category:      event.Category,
	}
}

func ToDomainEvent(mongoEvent MongoEvent) generic.Event {
	return generic.Event{
		Id:            mongoEvent.ID.Hex(),
		EventType:     mongoEvent.EventType,
		Payload:       mongoEvent.Payload,
		Timestamp:     mongoEvent.Timestamp,
		CorrelationId: mongoEvent.CorrelationID,
		Category:      mongoEvent.Category,
	}
}
