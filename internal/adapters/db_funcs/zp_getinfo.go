package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// CUTgetInfo calls the stored procedure 'dbo.CUTgetInfo()' on db.
func CUTgetInfo(ctx context.Context, db *sqlx.DB) error {
	// call dbo.CUTgetInfo
	const sqlstr = `dbo.CUTgetInfo`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr); err != nil {
		return logerror(err)
	}
	return nil
}
