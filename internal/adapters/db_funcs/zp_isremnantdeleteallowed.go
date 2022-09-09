package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTisRemnantDeleteAllowed calls the stored procedure 'dbo.CUTisRemnantDeleteAllowed(int, nvarchar)' on db.
func CUTisRemnantDeleteAllowed(ctx context.Context, db *sqlx.DB, iRemnantID int, strLang string) error {
	// call dbo.CUTisRemnantDeleteAllowed
	const sqlstr = `dbo.CUTisRemnantDeleteAllowed`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iRemnantID", iRemnantID), sql.Named("strLang", strLang)); err != nil {
		return logerror(err)
	}
	return nil
}
