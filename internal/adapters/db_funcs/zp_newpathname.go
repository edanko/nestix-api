package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// CUTnewPathName calls the stored procedure 'dbo.CUTnewPathName()' on db.
func CUTnewPathName(ctx context.Context, db *sqlx.DB) error {
	// call dbo.CUTnewPathName
	const sqlstr = `dbo.CUTnewPathName`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr); err != nil {
		return logerror(err)
	}
	return nil
}
