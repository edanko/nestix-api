package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// NxDbCharToReal calls the stored function 'dbo.NxDbCharToReal(nvarchar) real' on db.
func NxDbCharToReal(ctx context.Context, db *sqlx.DB, strText string) (float32, error) {
	// call dbo.NxDbCharToReal
	const sqlstr = `SELECT dbo.NxDbCharToReal(@p1) AS OUT`
	var r0 float32
	logf(tenantID, sqlstr, strText)
	if err := db.QueryRowContext(ctx, sqlstr, strText).Scan(&r0); err != nil {
		return 0.0, logerror(err)
	}
	return r0, nil
}
