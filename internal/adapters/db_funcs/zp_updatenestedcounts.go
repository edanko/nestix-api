package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTupdateNestedCounts calls the stored procedure 'dbo.CUTupdateNestedCounts(int, nvarchar)' on db.
func CUTupdateNestedCounts(ctx context.Context, db *sqlx.DB, iPathID int, strLang string) error {
	// call dbo.CUTupdateNestedCounts
	const sqlstr = `dbo.CUTupdateNestedCounts`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iPathID", iPathID), sql.Named("strLang", strLang)); err != nil {
		return logerror(err)
	}
	return nil
}
