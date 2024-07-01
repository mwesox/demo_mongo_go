package api

import (
	"cqrs_go/domain/generic"
	"cqrs_go/domain/product"
	"cqrs_go/domain/product/command"
	"cqrs_go/domain/product/event"
	"cqrs_go/domain/shared"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"time"
)

func PATCHProductUpdateDescription(c *gin.Context) {
	var changeProductDescription command.ChangeProductDescription
	if err := c.ShouldBindJSON(&changeProductDescription); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	eventRepo := shared.GetEventRepository(c)

	productDescriptionUpdatedEvent := event.ProductDescriptionUpdatedEvent{
		Description: changeProductDescription.Description,
	}

	payload, _ := json.Marshal(productDescriptionUpdatedEvent)

	productId := c.Param("id")

	//TODO handle error
	(*eventRepo).Save(generic.Event{
		CorrelationId: productId,
		EventType:     "productDescriptionUpdated",
		Timestamp:     time.Now(),
		Category:      "product",
		Payload:       payload,
	})

	c.JSON(200, gin.H{
		"productId": productId,
	})

}

func POSTProduct(c *gin.Context) {
	var addProduct command.AddProduct
	if err := c.BindJSON(&addProduct); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Println("Adding product: ", addProduct)

	repo, _ := c.Get("eventRepository")
	eventRepo, _ := repo.(generic.EventRepository)

	newProductId := uuid.NewString()

	productCreatedEvent := event.ProductCreatedEvent{
		ProductID:   newProductId,
		Description: addProduct.Description,
		Name:        addProduct.Name,
		Price:       addProduct.Price,
	}

	payload, _ := json.Marshal(productCreatedEvent)

	//TODO handle error
	eventRepo.Save(generic.Event{
		Payload:       payload,
		Category:      "product",
		EventType:     "productCreated",
		CorrelationId: newProductId,
	})

	c.JSON(200, gin.H{
		"productId": newProductId,
	})
}

func GETProduct(c *gin.Context) {
	id := c.Param("id")

	prd, _ := c.Get("productQueryService")
	productQueryService := prd.(product.ProductQueryService)

	product, err := productQueryService.FindById(id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, product)
}
