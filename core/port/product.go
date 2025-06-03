package port

import (
	"context"

	"github.com/WatkornFeng/go-hexa/core/domain"
	"github.com/WatkornFeng/go-hexa/core/domain/dto"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
}

type ProductService interface {
	CreateNewProduct(ctx context.Context, product *domain.Product) (*dto.ProductResponse, error)
}
