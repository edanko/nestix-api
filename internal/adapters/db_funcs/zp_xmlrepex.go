package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxXMLRepEx calls the stored procedure 'dbo.NxXmlRepEx(nvarchar, nvarchar)' on db.
func NxXMLRepEx(ctx context.Context, db *sqlx.DB, strProjectNo, strSection string) error {
	// call dbo.NxXmlRepEx
	const sqlstr = `dbo.NxXmlRepEx`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strProjectNo", strProjectNo), sql.Named("strSection", strSection)); err != nil {
		return logerror(err)
	}
	return nil
}
