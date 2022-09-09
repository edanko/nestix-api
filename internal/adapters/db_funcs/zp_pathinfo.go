package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTpathInfo calls the stored procedure 'dbo.CUTpathInfo(nvarchar, int, nvarchar)' on db.
func CUTpathInfo(ctx context.Context, db *sqlx.DB, strKeyword string, iID int, strLang string) error {
	// call dbo.CUTpathInfo
	const sqlstr = `dbo.CUTpathInfo`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strKeyword", strKeyword), sql.Named("iID", iID), sql.Named("strLang", strLang)); err != nil {
		return logerror(err)
	}
	return nil
}
