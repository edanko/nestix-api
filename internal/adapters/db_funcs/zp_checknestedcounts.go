package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTcheckNestedCounts calls the stored procedure 'dbo.CUTcheckNestedCounts(int, nvarchar)' on db.
func CUTcheckNestedCounts(ctx context.Context, db *sqlx.DB, iPathID int, strLang string) error {
	// call dbo.CUTcheckNestedCounts
	const sqlstr = `dbo.CUTcheckNestedCounts`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iPathID", iPathID), sql.Named("strLang", strLang)); err != nil {
		return logerror(err)
	}
	return nil
}
