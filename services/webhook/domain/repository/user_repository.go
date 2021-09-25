package repository

import (
	"context"
	"webhook/domain/model"
)

type UserRepository interface {
	FindById(ctx context.Context, id string) (*model.User, error)
	// FindAll(ctx context.Context) ([]model.User, error)
	Store(ctx context.Context, user *model.User) error
	// Update(ctx context.Context, user *model.User) error
	// Delete(ctx context.Context, user *model.User) error
}
