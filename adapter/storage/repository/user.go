package repository

import (
	"github.com/WatkornFeng/go-hexa/core/domain"
	"github.com/WatkornFeng/go-hexa/core/port"
	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &userRepositoryDB{db: db}
}

func (r *userRepositoryDB) GetAll() ([]domain.User, error) {
	var users []domain.User
	if result := r.db.Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
