package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTafterUpdateOrderline calls the stored procedure 'dbo.CUTafterUpdateOrderline(int, nvarchar)' on db.
func CUTafterUpdateOrderline(ctx context.Context, db *sqlx.DB, iPosID int, strChanger string) error {
	// call dbo.CUTafterUpdateOrderline
	const sqlstr = `dbo.CUTafterUpdateOrderline`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iPosID", iPosID), sql.Named("strChanger", strChanger)); err != nil {
		return logerror(err)
	}
	return nil
}
