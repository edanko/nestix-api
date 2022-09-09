package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTdeletePathParts calls the stored procedure 'dbo.CUTdeletePathParts(int, nvarchar)' on db.
func CUTdeletePathParts(ctx context.Context, db *sqlx.DB, iPathID int, strLang string) error {
	// call dbo.CUTdeletePathParts
	const sqlstr = `dbo.CUTdeletePathParts`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iPathID", iPathID), sql.Named("strLang", strLang)); err != nil {
		return logerror(err)
	}
	return nil
}
