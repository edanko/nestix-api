package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTcreateUnbornRemnant calls the stored procedure 'dbo.CUTcreateUnbornRemnant(int)' on db.
func CUTcreateUnbornRemnant(ctx context.Context, db *sqlx.DB, iSheetPathDetID int) error {
	// call dbo.CUTcreateUnbornRemnant
	const sqlstr = `dbo.CUTcreateUnbornRemnant`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iSheetPathDetID", iSheetPathDetID)); err != nil {
		return logerror(err)
	}
	return nil
}
