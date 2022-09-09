package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// NxDbGetPartCutLength calls the stored function 'dbo.NxDbGetPartCutLength(int) int' on db.
func NxDbGetPartCutLength(ctx context.Context, db *sqlx.DB, iPartID int) (int, error) {
	// call dbo.NxDbGetPartCutLength
	const sqlstr = `SELECT dbo.NxDbGetPartCutLength(@p1) AS OUT`
	var r0 int
	logf(tenantID, sqlstr, iPartID)
	if err := db.QueryRowContext(ctx, sqlstr, iPartID).Scan(&r0); err != nil {
		return 0, logerror(err)
	}
	return r0, nil
}
