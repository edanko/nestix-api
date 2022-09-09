package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxMaintCheckEmptyColumns calls the stored procedure 'dbo.nx_maint_checkEmptyColumns(int)' on db.
func NxMaintCheckEmptyColumns(ctx context.Context, db *sqlx.DB, iInitTables int) error {
	// call dbo.nx_maint_checkEmptyColumns
	const sqlstr = `dbo.nx_maint_checkEmptyColumns`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iInitTables", iInitTables)); err != nil {
		return logerror(err)
	}
	return nil
}
