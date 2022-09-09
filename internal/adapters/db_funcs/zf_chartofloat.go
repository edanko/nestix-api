package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// NxDbCharToFloat calls the stored function 'dbo.NxDbCharToFloat(nvarchar) float' on db.
func NxDbCharToFloat(ctx context.Context, db *sqlx.DB, strText string) (float64, error) {
	// call dbo.NxDbCharToFloat
	const sqlstr = `SELECT dbo.NxDbCharToFloat(@p1) AS OUT`
	var r0 float64
	logf(tenantID, sqlstr, strText)
	if err := db.QueryRowContext(ctx, sqlstr, strText).Scan(&r0); err != nil {
		return 0.0, logerror(err)
	}
	return r0, nil
}
