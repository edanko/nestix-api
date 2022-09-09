package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// NxDbGetPartLineMarkingLength calls the stored function 'dbo.NxDbGetPartLineMarkingLength(int) float' on db.
func NxDbGetPartLineMarkingLength(ctx context.Context, db *sqlx.DB, iPartID int) (float64, error) {
	// call dbo.NxDbGetPartLineMarkingLength
	const sqlstr = `SELECT dbo.NxDbGetPartLineMarkingLength(@p1) AS OUT`
	var r0 float64
	logf(tenantID, sqlstr, iPartID)
	if err := db.QueryRowContext(ctx, sqlstr, iPartID).Scan(&r0); err != nil {
		return 0.0, logerror(err)
	}
	return r0, nil
}
