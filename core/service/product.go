package service

import (
	"context"
	"errors"
	"log/slog"
	"math"
	"time"

	"github.com/WatkornFeng/go-hexa/core/domain"
	"github.com/WatkornFeng/go-hexa/core/domain/dto"
	"github.com/WatkornFeng/go-hexa/core/port"
)

type productService struct {
	repo  port.ProductRepository
	cache port.CacheRepository
}

func NewProductService(repo port.ProductRepository, cache port.CacheRepository) port.ProductService {
	return &productService{repo, cache}
}

func (us *productService) CreateNewProduct(ctx context.Context, product *domain.Product) (*dto.ProductResponse, error) {
	// validate price should not have more than 2 decimals
	if !(math.Round(product.Price*100) == product.Price*100) {
		return nil, domain.ErrProductPriceNotCorrect
	}
	createCtx, cancelCreate := context.WithTimeout(ctx, 1*time.Second)
	defer cancelCreate()
	product, err := us.repo.CreateProduct(createCtx, product)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, domain.ErrDatabaseTimeOut

		}
		slog.Error("Error creating product", "error", err)
		return nil, domain.ErrInternalServerError
	}
	rspProduct := dto.NewProductResponse(product)
	return rspProduct, nil
}
