package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTnewPathNameShip2 calls the stored procedure 'dbo.CUTnewPathNameShip_2(nvarchar)' on db.
func CUTnewPathNameShip2(ctx context.Context, db *sqlx.DB, strOrderlineid string) error {
	// call dbo.CUTnewPathNameShip_2
	const sqlstr = `dbo.CUTnewPathNameShip_2`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strOrderlineid", strOrderlineid)); err != nil {
		return logerror(err)
	}
	return nil
}
