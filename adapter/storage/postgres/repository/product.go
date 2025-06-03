package repository

import (
	"context"

	"github.com/WatkornFeng/go-hexa/core/domain"
	"github.com/WatkornFeng/go-hexa/core/port"
	"gorm.io/gorm"
)

type productRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) port.ProductRepository {
	return &productRepositoryDB{db: db}
}

func (r *productRepositoryDB) CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	if err := r.db.WithContext(ctx).Create(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
