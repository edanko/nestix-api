package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// NxDbGetPartPointMarkingCount calls the stored function 'dbo.NxDbGetPartPointMarkingCount(int) int' on db.
func NxDbGetPartPointMarkingCount(ctx context.Context, db *sqlx.DB, iPartID int) (int, error) {
	// call dbo.NxDbGetPartPointMarkingCount
	const sqlstr = `SELECT dbo.NxDbGetPartPointMarkingCount(@p1) AS OUT`
	var r0 int
	logf(tenantID, sqlstr, iPartID)
	if err := db.QueryRowContext(ctx, sqlstr, iPartID).Scan(&r0); err != nil {
		return 0, logerror(err)
	}
	return r0, nil
}
