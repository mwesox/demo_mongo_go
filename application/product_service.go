package application

import (
	"cqrs_go/domain/generic"
)

const CATEGORY = "product"

type ProductService struct {
	repo generic.EventRepository
}

func NewProductService(repo generic.EventRepository) *ProductService {
	return &ProductService{repo: repo}
}

/*func (p *ProductService) GetAllProducts() ([]product.Product, error) {
	events, err := p.repo.FindByCategory(CATEGORY)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return events, nil
}*/

/*func (p *ProductService) GetProductByID(id string) ([]product.Product, error) {
	// This function should consume events with correlationId = id
	events, err := p.repo.FindByCorrelationID(id)
	if err != nil {
		return nil, err
	}
	return events, nil
}*/
