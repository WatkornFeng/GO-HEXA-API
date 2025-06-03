package dto

import "github.com/WatkornFeng/go-hexa/core/domain"

// dto => Data Transfer Object
type ProductResponse struct {
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
}

func NewProductResponse(product *domain.Product) *ProductResponse {
	return &ProductResponse{
		ProductID: product.ID,
		Name:      product.Name,
		Price:     product.Price,
	}
}
