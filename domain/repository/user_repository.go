package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/kagizi/kc-tech-test/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetByPhone(ctx context.Context, phone string) (*entity.User, error)
}
