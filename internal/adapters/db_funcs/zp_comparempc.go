package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTcompareMPC calls the stored procedure 'dbo.CUTcompareMPC (nvarchar, nvarchar)' on db.
func CUTcompareMPC(ctx context.Context, db *sqlx.DB, strOperation, strMPCset string) error {
	// call dbo.CUTcompareMPC
	const sqlstr = `dbo.CUTcompareMPC `
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strOperation", strOperation), sql.Named("strMPCset", strMPCset)); err != nil {
		return logerror(err)
	}
	return nil
}
