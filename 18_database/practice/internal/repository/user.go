package repository

import (
	"context"
	"errors"

	"go3hours/18_database/practice/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("user nof found")

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) Create(ctx context.Context, u model.UserRequest) (model.User, error) {
	var user model.User
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email, created_at`
	err := r.pool.QueryRow(ctx, query, u.Name, u.Email).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	return user, err
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (model.User, error) {
	var user model.User
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1`
	err := r.pool.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return model.User{}, ErrNotFound
	}

	return user, err
}

func (r *UserRepository) List(ctx context.Context) ([]model.User, error) {
	query := `SELECT id, name, email, created_at FROM users ORDER BY id`

	rows, err := r.pool.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func (r *UserRepository) Update(ctx context.Context, id int, u model.UserRequest) (model.User, error) {
	var user model.User

	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3 RETURNING id, name, email, created_at`

	err := r.pool.QueryRow(ctx, query, u.Name, u.Email, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return model.User{}, ErrNotFound
	}
	return user, err
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	tag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
