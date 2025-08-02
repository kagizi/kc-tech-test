package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/kagizi/kc-tech-test/domain/entity"
	"github.com/kagizi/kc-tech-test/domain/repository"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	query := `
        INSERT INTO users (id, name, phone, created_at, updated_at)
        VALUES (UUID_TO_BIN(?), ?, ?, ?, ?)
    `
	_, err := r.db.ExecContext(ctx, query,
		user.ID.String(),
		user.Name,
		user.Phone,
		user.CreatedAt,
		user.UpdatedAt,
	)
	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	query := `
        SELECT BIN_TO_UUID(id), name, phone, created_at, updated_at
        FROM users
        WHERE id = UUID_TO_BIN(?)
    `
	row := r.db.QueryRowContext(ctx, query, id.String())

	var user entity.User
	var idStr string
	err := row.Scan(
		&idStr,
		&user.Name,
		&user.Phone,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	user.ID, err = uuid.Parse(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) GetByPhone(ctx context.Context, phone string) (*entity.User, error) {
	query := `
        SELECT BIN_TO_UUID(id), name, phone, created_at, updated_at
        FROM users
        WHERE phone = ?
    `
	row := r.db.QueryRowContext(ctx, query, phone)

	var user entity.User
	var idStr string
	err := row.Scan(
		&idStr,
		&user.Name,
		&user.Phone,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	user.ID, err = uuid.Parse(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %w", err)
	}

	return &user, nil
}
