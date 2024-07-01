package shared

import (
	"cqrs_go/domain/generic"
	"cqrs_go/domain/product"
	"errors"
	"github.com/gin-gonic/gin"
)

const EVENT_REPOSITORY = "eventRepository"
const PRODUCT_QUERY_SERVICE = "productQueryService"

func SetEventRepository(c *gin.Context, eventRepository *generic.EventRepository) {
	(*c).Set(EVENT_REPOSITORY, eventRepository)
}

func GetEventRepository(c *gin.Context) *generic.EventRepository {
	repo, exists := (*c).Get(EVENT_REPOSITORY)

	if !exists {
		panic(errors.New("Event repository doesnt exist"))
	}

	eventRepository := repo.(generic.EventRepository)

	return &eventRepository

}

func SetProductQueryService(c *gin.Context, service *product.ProductQueryService) {
	(*c).Set(PRODUCT_QUERY_SERVICE, service)
}

func GetProductQueryService(c *gin.Context) *product.ProductQueryService {
	s, exists := (*c).Get(PRODUCT_QUERY_SERVICE)
	if !exists {
		panic(errors.New("Product query service doesnt exist"))
	}

	service := s.(product.ProductQueryService)

	return &service

}
