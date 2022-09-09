package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTcanDeletePath calls the stored procedure 'dbo.CUTcanDeletePath(int, nvarchar)' on db.
func CUTcanDeletePath(ctx context.Context, db *sqlx.DB, iPathID int, strLang string) error {
	// call dbo.CUTcanDeletePath
	const sqlstr = `dbo.CUTcanDeletePath`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iPathID", iPathID), sql.Named("strLang", strLang)); err != nil {
		return logerror(err)
	}
	return nil
}
