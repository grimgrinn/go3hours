package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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
