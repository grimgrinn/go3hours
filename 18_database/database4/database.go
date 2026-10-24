package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upSeedRoles, downSeedRoles)
}

func upSeedRoles(ctx context.Context, tx *sql.Tx) error {
	roles := []string{"admin", "moderator", "user"}
	for _, role := range roles {
		if _, err := tx.ExecContext(ctx,
			"INSERT INTO roles (name) VALUES ($1) ON CONFLICT DO NOTHING", role,
		); err != nil {
			return err
		}
	}
	return nil
}

func downSeedRoles(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DELETE FROM roles WHERE name IN ('admin', 'moderator', 'user')")
	return err
}
