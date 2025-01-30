package userrepository

import (
	"context"
	"mindscribe-be/internal/models"
	baserepository "mindscribe-be/internal/repository/base-repository"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	base *baserepository.BaseRepository
}

func NewUserRepo(base *baserepository.BaseRepository) *UserRepo {
	return &UserRepo{
		base: base,
	}
}

func (r UserRepo) StoreUser(ctx context.Context, exec sqlx.ExtContext, user *models.User) error {
	query := `
        INSERT INTO users (email, username, password, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
	row := exec.QueryRowxContext(ctx, query, user.Email, user.Username, user.Password, user.CreatedAt, user.UpdatedAt)
	if err := row.Scan(&user.ID); err != nil {
		return err
	}
	return nil
}
