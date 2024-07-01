package product

import (
	"cqrs_go/domain/generic"
	pevents "cqrs_go/domain/product/event"
	"encoding/json"
)

type ProductQueryService struct {
	repo *generic.EventRepository
}

func NewProductQueryService(repo *generic.EventRepository) ProductQueryService {
	return ProductQueryService{repo: repo}
}

func (service *ProductQueryService) FindById(productId string) (*Product, error) {
	events, _ := (*service.repo).FindByCorrelationID(productId)

	product := &Product{}

	for _, event := range *events {
		apply(product, event)
	}

	return product, nil

}

func apply(product *Product, event generic.Event) {

	switch event.EventType {
	case "productCreated":
		var ev pevents.ProductCreatedEvent
		json.Unmarshal(event.Payload, &ev)
		product.ProductID = ev.ProductID
		product.Name = ev.Name
		product.Description = ev.Description

	case "productDescriptionUpdated":
		var ev pevents.ProductDescriptionUpdatedEvent
		json.Unmarshal(event.Payload, &ev)
		product.Description = ev.Description
	}
}
