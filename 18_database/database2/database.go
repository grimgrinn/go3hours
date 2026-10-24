package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID        int64
	FullName  string
	Email     string
	CreatedAt time.Time
}

type postgresUserRepo struct {
	pool *pgxpool.Pool
}

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

func mapDBError(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // unique_violation
			return ErrAlreadyExists
		case "23503": // foreign_key_violation
			return ErrNotFound
		}
	}
	return err
}

func (r *postgresUserRepo) GetByID(ctx context.Context, id int64) (User, error) {
	var u User

	err := r.pool.QueryRow(ctx,
		"SELECT id, full_name, email, created_at FROM users WHERE id = $1",
		id,
	).Scan(&u.ID, &u.FullName, &u.Email, &u.CreatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return User{}, ErrNotFound
	}
	return u, err

}

func (r *postgresUserRepo) List(ctx context.Context, limit, offset int) ([]User, error) {
	rows, err := r.pool.Query(ctx,
		"SELECT id, full_name, email, create_at FROM users ORDER BY id LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.FullName, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()

}

func (r *postgresUserRepo) Create(ctx context.Context, name, email string) (User, error) {
	var u User

	err := r.pool.QueryRow(ctx,
		"INSERT INTO users (full_name, email) VALUES ($1, $2) RETURNING id, full_name, email, created_at",
		name, email).Scan(&u.ID, &u.FullName, &u.Email, &u.CreatedAt)

	return u, err
}

func main() {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	repo := &postgresUserRepo{pool: pool}
	ctx := context.Background()

	u, err := repo.Create(ctx, "Adolf Hitler", "adolfh@zz.de")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
}
