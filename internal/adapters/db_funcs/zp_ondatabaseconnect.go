package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTonDatabaseConnect calls the stored procedure 'dbo.CUTonDatabaseConnect(nvarchar, nvarchar)' on db.
func CUTonDatabaseConnect(ctx context.Context, db *sqlx.DB, strSite, strUser string) error {
	// call dbo.CUTonDatabaseConnect
	const sqlstr = `dbo.CUTonDatabaseConnect`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strSite", strSite), sql.Named("strUser", strUser)); err != nil {
		return logerror(err)
	}
	return nil
}
