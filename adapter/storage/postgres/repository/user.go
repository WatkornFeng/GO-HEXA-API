package repository

import (
	"context"
	"errors"

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

func (r *userRepositoryDB) GetAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User

	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepositoryDB) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepositoryDB) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Not found is not an error, let service decide
		}
		return nil, err
	}
	return &user, nil

}

func (r *userRepositoryDB) GetUserByID(ctx context.Context, id uint64) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).Where("id = ?", uint(id)).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Not found is not an error, let service decide
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryDB) UpdateUserByID(ctx context.Context, id uint64, updateData *domain.User) (*domain.User, error) {

	// 1. Update
	result := r.db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", uint(id)).Update("name", updateData.Name)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}
	// 2. Fetch updated user
	var updatedUser domain.User
	if err := r.db.WithContext(ctx).First(&updatedUser, uint(id)).Error; err != nil {
		return nil, err
	}

	return &updatedUser, nil
}
