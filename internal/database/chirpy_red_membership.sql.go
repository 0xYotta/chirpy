// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: chirpy_red_membership.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const upgradeUserToRed = `-- name: UpgradeUserToRed :exec
UPDATE users
SET
    is_chirpy_red = TRUE
WHERE id = $1
`

func (q *Queries) UpgradeUserToRed(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, upgradeUserToRed, id)
	return err
}
