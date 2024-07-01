package product

import (
	"cqrs_go/domain/generic"
	"cqrs_go/domain/product/_tests"
	pevents "cqrs_go/domain/product/event"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindById(t *testing.T) {
	// Prepare the mock data
	productCreatedEvent := pevents.ProductCreatedEvent{
		ProductID:   "1",
		Name:        "Test Product",
		Description: "Test Description",
	}
	productCreatedPayload, _ := json.Marshal(productCreatedEvent)

	productDescriptionUpdatedEvent := pevents.ProductDescriptionUpdatedEvent{
		Description: "Updated Description",
	}
	productDescriptionUpdatedPayload, _ := json.Marshal(productDescriptionUpdatedEvent)

	events := []generic.Event{
		{EventType: "productCreated", Payload: productCreatedPayload},
		{EventType: "productDescriptionUpdated", Payload: productDescriptionUpdatedPayload},
	}

	// Create a mock repository with events mapped by correlation ID
	var mockRepo generic.EventRepository
	mockRepo = &_tests.MockEventRepository{
		EventsByCorrelationID: map[string][]generic.Event{
			"1": events,
		},
	}

	// Create the service with the mock repository
	service := NewProductQueryService(&mockRepo)

	// Call the service
	product, err := service.FindById("1")

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "1", product.ProductID)
	assert.Equal(t, "Test Product", product.Name)
	assert.Equal(t, "Updated Description", product.Description)
}
