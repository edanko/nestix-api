package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// NxDbGetRemnantsDimsForNesting calls the stored function 'dbo.NxDbGetRemnantsDimsForNesting(int) nvarchar' on db.
func NxDbGetRemnantsDimsForNesting(ctx context.Context, db *sqlx.DB, iPathID int) (string, error) {
	// call dbo.NxDbGetRemnantsDimsForNesting
	const sqlstr = `SELECT dbo.NxDbGetRemnantsDimsForNesting(@p1) AS OUT`
	var r0 string
	logf(tenantID, sqlstr, iPathID)
	if err := db.QueryRowContext(ctx, sqlstr, iPathID).Scan(&r0); err != nil {
		return "", logerror(err)
	}
	return r0, nil
}
