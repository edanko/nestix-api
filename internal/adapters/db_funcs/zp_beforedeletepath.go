package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTbeforeDeletePath calls the stored procedure 'dbo.CUTbeforeDeletePath(int, nvarchar)' on db.
func CUTbeforeDeletePath(ctx context.Context, db *sqlx.DB, iPathID int, strLang string) error {
	// call dbo.CUTbeforeDeletePath
	const sqlstr = `dbo.CUTbeforeDeletePath`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iPathID", iPathID), sql.Named("strLang", strLang)); err != nil {
		return logerror(err)
	}
	return nil
}
