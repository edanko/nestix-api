package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTpartInfo calls the stored procedure 'dbo.CUTpartInfo(nvarchar, int, nvarchar)' on db.
func CUTpartInfo(ctx context.Context, db *sqlx.DB, strKeyWord string, iID int, strLang string) error {
	// call dbo.CUTpartInfo
	const sqlstr = `dbo.CUTpartInfo`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strKeyWord", strKeyWord), sql.Named("iID", iID), sql.Named("strLang", strLang)); err != nil {
		return logerror(err)
	}
	return nil
}
