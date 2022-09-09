package db_funcs

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

// NxDbCharToDateTime calls the stored function 'dbo.NxDbCharToDateTime(nvarchar) datetime' on db.
func NxDbCharToDateTime(ctx context.Context, db *sqlx.DB, strText string) (time.Time, error) {
	// call dbo.NxDbCharToDateTime
	const sqlstr = `SELECT dbo.NxDbCharToDateTime(@p1) AS OUT`
	var r0 time.Time
	logf(tenantID, sqlstr, strText)
	if err := db.QueryRowContext(ctx, sqlstr, strText).Scan(&r0); err != nil {
		return time.Time{}, logerror(err)
	}
	return r0, nil
}
