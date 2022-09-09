package db_funcs

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

// NxDbDateTimeToUTCString calls the stored function 'dbo.NxDbDateTimeToUTCString(datetime) nvarchar' on db.
func NxDbDateTimeToUTCString(ctx context.Context, db *sqlx.DB, dtDateTime time.Time) (string, error) {
	// call dbo.NxDbDateTimeToUTCString
	const sqlstr = `SELECT dbo.NxDbDateTimeToUTCString(@p1) AS OUT`
	var r0 string
	logf(tenantID, sqlstr, dtDateTime)
	if err := db.QueryRowContext(ctx, sqlstr, dtDateTime).Scan(&r0); err != nil {
		return "", logerror(err)
	}
	return r0, nil
}
