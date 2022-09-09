package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// CUTnewRemnantName calls the stored procedure 'dbo.CUTnewRemnantName()' on db.
func CUTnewRemnantName(ctx context.Context, db *sqlx.DB) error {
	// call dbo.CUTnewRemnantName
	const sqlstr = `dbo.CUTnewRemnantName`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr); err != nil {
		return logerror(err)
	}
	return nil
}
