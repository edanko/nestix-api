package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// NxDbGetConfigOption calls the stored function 'dbo.NxDbGetConfigOption(nvarchar) int' on db.
func NxDbGetConfigOption(ctx context.Context, db *sqlx.DB, strOptionName string) (int, error) {
	// call dbo.NxDbGetConfigOption
	const sqlstr = `SELECT dbo.NxDbGetConfigOption(@p1) AS OUT`
	var r0 int
	logf(tenantID, sqlstr, strOptionName)
	if err := db.QueryRowContext(ctx, sqlstr, strOptionName).Scan(&r0); err != nil {
		return 0, logerror(err)
	}
	return r0, nil
}
