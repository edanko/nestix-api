package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTisNestDeleteAllowed calls the stored procedure 'dbo.CUTisNestDeleteAllowed(int, nvarchar, nvarchar)' on db.
func CUTisNestDeleteAllowed(
	ctx context.Context,
	db *sqlx.DB,
	iPathID int,
	strUser, strLang string,
) error {
	// call dbo.CUTisNestDeleteAllowed
	const sqlstr = `dbo.CUTisNestDeleteAllowed`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iPathID", iPathID), sql.Named("strUser", strUser), sql.Named("strLang", strLang)); err != nil {
		return logerror(err)
	}
	return nil
}
