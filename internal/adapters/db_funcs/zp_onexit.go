package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// CUTonExit calls the stored procedure 'dbo.CUTonExit()' on db.
func CUTonExit(ctx context.Context, db *sqlx.DB) error {
	// call dbo.CUTonExit
	const sqlstr = `dbo.CUTonExit`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr); err != nil {
		return logerror(err)
	}
	return nil
}
