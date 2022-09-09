package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTupdateMachine calls the stored procedure 'dbo.CUTupdateMachine(nvarchar)' on db.
func CUTupdateMachine(ctx context.Context, db *sqlx.DB, strMachine string) error {
	// call dbo.CUTupdateMachine
	const sqlstr = `dbo.CUTupdateMachine`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strMachine", strMachine)); err != nil {
		return logerror(err)
	}
	return nil
}
