package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTqueryInfo calls the stored procedure 'dbo.CUTqueryInfo(xml)' on db.
func CUTqueryInfo(ctx context.Context, db *sqlx.DB, xmlData []byte) error {
	// call dbo.CUTqueryInfo
	const sqlstr = `dbo.CUTqueryInfo`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("xmlData", xmlData)); err != nil {
		return logerror(err)
	}
	return nil
}
