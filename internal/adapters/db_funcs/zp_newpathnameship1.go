package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTnewPathNameShip1 calls the stored procedure 'dbo.CUTnewPathNameShip_1(nvarchar)' on db.
func CUTnewPathNameShip1(ctx context.Context, db *sqlx.DB, strOrderlineid string) error {
	// call dbo.CUTnewPathNameShip_1
	const sqlstr = `dbo.CUTnewPathNameShip_1`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strOrderlineid", strOrderlineid)); err != nil {
		return logerror(err)
	}
	return nil
}
