package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTorderInfo calls the stored procedure 'dbo.CUTorderInfo(nvarchar, int, nvarchar)' on db.
func CUTorderInfo(ctx context.Context, db *sqlx.DB, strKeyWord string, iID int, strLang string) error {
	// call dbo.CUTorderInfo
	const sqlstr = `dbo.CUTorderInfo`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strKeyWord", strKeyWord), sql.Named("iID", iID), sql.Named("strLang", strLang)); err != nil {
		return logerror(err)
	}
	return nil
}
