package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// NxDbCharToInt calls the stored function 'dbo.NxDbCharToInt(nvarchar) int' on db.
func NxDbCharToInt(ctx context.Context, db *sqlx.DB, strText string) (int, error) {
	// call dbo.NxDbCharToInt
	const sqlstr = `SELECT dbo.NxDbCharToInt(@p1) AS OUT`
	var r0 int
	logf(tenantID, sqlstr, strText)
	if err := db.QueryRowContext(ctx, sqlstr, strText).Scan(&r0); err != nil {
		return 0, logerror(err)
	}
	return r0, nil
}
