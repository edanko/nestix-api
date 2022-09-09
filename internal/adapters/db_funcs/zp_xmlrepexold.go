package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxXMLRepExOld calls the stored procedure 'dbo.NxXmlRepEx_old(nvarchar, nvarchar)' on db.
func NxXMLRepExOld(ctx context.Context, db *sqlx.DB, strProject, strSection string) error {
	// call dbo.NxXmlRepEx_old
	const sqlstr = `dbo.NxXmlRepEx_old`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strProject", strProject), sql.Named("strSection", strSection)); err != nil {
		return logerror(err)
	}
	return nil
}
