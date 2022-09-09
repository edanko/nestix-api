package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTafterPartImportProc calls the stored procedure 'dbo.CUTafterPartImportProc(int, nvarchar)' on db.
func CUTafterPartImportProc(ctx context.Context, db *sqlx.DB, iPartID int, strSetName string) error {
	// call dbo.CUTafterPartImportProc
	const sqlstr = `dbo.CUTafterPartImportProc`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iPartID", iPartID), sql.Named("strSetName", strSetName)); err != nil {
		return logerror(err)
	}
	return nil
}
