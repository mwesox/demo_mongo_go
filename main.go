package main

import (
	"cqrs_go/api"
	"cqrs_go/domain/generic"
	"cqrs_go/domain/product"
	"cqrs_go/domain/shared"
	"cqrs_go/infrastructure/mongodb"
	"github.com/gin-gonic/gin"
)

var eventRepository generic.EventRepository
var productQueryService product.ProductQueryService

func init() {
	eventRepository = mongodb.NewRepository()
	productQueryService = product.NewProductQueryService(&eventRepository)
}

func main() {

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(func(context *gin.Context) {
		shared.SetEventRepository(context, &eventRepository)
		shared.SetProductQueryService(context, &productQueryService)
		context.Next()
	})

	r.POST("/products", api.POSTProduct)
	r.GET("/products/:id", api.GETProduct)
	r.PATCH("/products/:id", api.PATCHProductUpdateDescription)

	r.Run()
}
