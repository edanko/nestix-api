package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// CUTnewSheetName calls the stored procedure 'dbo.CUTnewSheetName()' on db.
func CUTnewSheetName(ctx context.Context, db *sqlx.DB) error {
	// call dbo.CUTnewSheetName
	const sqlstr = `dbo.CUTnewSheetName`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr); err != nil {
		return logerror(err)
	}
	return nil
}
