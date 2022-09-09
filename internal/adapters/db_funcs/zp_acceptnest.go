package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTacceptNest calls the stored procedure 'dbo.CUTacceptNest(int)' on db.
func CUTacceptNest(ctx context.Context, db *sqlx.DB, iPathID int) error {
	// call dbo.CUTacceptNest
	const sqlstr = `dbo.CUTacceptNest`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iPathID", iPathID)); err != nil {
		return logerror(err)
	}
	return nil
}
