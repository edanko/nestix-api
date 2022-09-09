package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTnewPathNameFromTemplate calls the stored procedure 'dbo.CUTnewPathNameFromTemplate(nvarchar)' on db.
func CUTnewPathNameFromTemplate(ctx context.Context, db *sqlx.DB, strOrderlineid string) error {
	// call dbo.CUTnewPathNameFromTemplate
	const sqlstr = `dbo.CUTnewPathNameFromTemplate`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strOrderlineid", strOrderlineid)); err != nil {
		return logerror(err)
	}
	return nil
}
