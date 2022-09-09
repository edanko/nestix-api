package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// NxDbGetPartCutStartCount calls the stored function 'dbo.NxDbGetPartCutStartCount(int) int' on db.
func NxDbGetPartCutStartCount(ctx context.Context, db *sqlx.DB, iPartID int) (int, error) {
	// call dbo.NxDbGetPartCutStartCount
	const sqlstr = `SELECT dbo.NxDbGetPartCutStartCount(@p1) AS OUT`
	var r0 int
	logf(tenantID, sqlstr, iPartID)
	if err := db.QueryRowContext(ctx, sqlstr, iPartID).Scan(&r0); err != nil {
		return 0, logerror(err)
	}
	return r0, nil
}
