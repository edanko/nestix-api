package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// NxDbGetRemnantParentNestName calls the stored function 'dbo.NxDbGetRemnantParentNestName(int) nvarchar' on db.
func NxDbGetRemnantParentNestName(ctx context.Context, db *sqlx.DB, iRemnantID int) (string, error) {
	// call dbo.NxDbGetRemnantParentNestName
	const sqlstr = `SELECT dbo.NxDbGetRemnantParentNestName(@p1) AS OUT`
	var r0 string
	logf(tenantID, sqlstr, iRemnantID)
	if err := db.QueryRowContext(ctx, sqlstr, iRemnantID).Scan(&r0); err != nil {
		return "", logerror(err)
	}
	return r0, nil
}
