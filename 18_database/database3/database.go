package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepo struct {
	pool *pgxpool.Pool
}

func (r *postgresRepo) TransferMoney(ctx context.Context, fromID, toID int64, amount float64) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx,
		"UPDATE accounts SET balance = balance - $1 WHERE id = $2",
		amount, fromID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx,
		"UPDATE accounts SET balance + $1 WHERE id = $2",
		amount, toID,
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
