package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTupdatePathData calls the stored procedure 'dbo.CUTupdatePathData(int)' on db.
func CUTupdatePathData(ctx context.Context, db *sqlx.DB, iPathID int) error {
	// call dbo.CUTupdatePathData
	const sqlstr = `dbo.CUTupdatePathData`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iPathID", iPathID)); err != nil {
		return logerror(err)
	}
	return nil
}
