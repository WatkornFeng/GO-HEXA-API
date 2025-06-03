package service

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/WatkornFeng/go-hexa/core/domain"
	"github.com/WatkornFeng/go-hexa/core/domain/dto"
	"github.com/WatkornFeng/go-hexa/core/port"
	"github.com/WatkornFeng/go-hexa/core/util"
)

type userService struct {
	repo  port.UserRepository
	cache port.CacheRepository
}

func NewUserService(repo port.UserRepository, cache port.CacheRepository) port.UserService {
	return &userService{repo, cache}
}

func (us *userService) GetUsers(ctx context.Context) ([]dto.UserResponse, error) {
	var rspListUsers []dto.UserResponse
	// Redis GET
	cacheKey := util.GenerateCacheKey("users", "all")
	cachedUsers, err := us.cache.Get(ctx, cacheKey)
	if err == nil {
		err := util.Deserialize(cachedUsers, &rspListUsers)
		if err == nil {
			slog.Info("GET DAT FROM REDIS CACHE")
			return rspListUsers, nil
		}
	}

	// Repository
	dbCtx, cancelDB := context.WithTimeout(ctx, 1*time.Second)
	defer cancelDB()

	// time.Sleep(3 * time.Second)
	users, err := us.repo.GetAll(dbCtx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, domain.ErrDatabaseTimeOut
		}
		slog.Error("Error retrieving users", "error", err)
		return nil, domain.ErrInternalServerError
	}
	rspListUsers = dto.NewListUsersResponse(users)

	// Redis SET
	usersSerialized, err := util.Serialize(rspListUsers)
	if err == nil {
		us.cache.Set(ctx, cacheKey, usersSerialized, time.Second*60)
	}

	slog.Info("GET DAT FROM DB")
	return rspListUsers, nil
}

func (us *userService) Register(ctx context.Context, user *domain.User) (*dto.UserResponse, error) {
	// Check if email already exist
	findCtx, cancelFind := context.WithTimeout(ctx, 1*time.Second)
	defer cancelFind()
	existing, err := us.repo.FindUserByEmail(findCtx, user.Email)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, domain.ErrDatabaseTimeOut

		}
		slog.Error("Error Finding user by email", "error", err)
		return nil, domain.ErrInternalServerError
	}
	if existing != nil {
		// Business rule: cannot register if email exists
		return nil, domain.ErrUserAlreadyExists
	}

	// if not exist then create new one
	createCtx, cancelCreate := context.WithTimeout(ctx, 1*time.Second)
	defer cancelCreate()
	user, err = us.repo.CreateUser(createCtx, user)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, domain.ErrDatabaseTimeOut

		}
		slog.Error("Error creating user", "error", err)
		return nil, domain.ErrInternalServerError
	}
	rspUser := dto.NewUserResponse(user)

	return rspUser, nil
}
