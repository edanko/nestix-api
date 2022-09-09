package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTafterOrderListImport calls the stored procedure 'dbo.CUTafterOrderListImport(int)' on db.
func CUTafterOrderListImport(ctx context.Context, db *sqlx.DB, iPartID int) error {
	// call dbo.CUTafterOrderListImport
	const sqlstr = `dbo.CUTafterOrderListImport`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iPartID", iPartID)); err != nil {
		return logerror(err)
	}
	return nil
}
