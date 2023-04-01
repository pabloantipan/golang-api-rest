package repository

import (
	"context"
	"go/golang-api-rest/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id string) (*models.User, error)
	Close() error
}

var implementation UserRepository

func SetRepository(repository UserRepository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUser(ctx, id)
}
