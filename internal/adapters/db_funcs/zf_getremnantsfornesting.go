package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// NxDbGetRemnantsForNesting calls the stored function 'dbo.NxDbGetRemnantsForNesting(int) nvarchar' on db.
func NxDbGetRemnantsForNesting(ctx context.Context, db *sqlx.DB, iPathID int) (string, error) {
	// call dbo.NxDbGetRemnantsForNesting
	const sqlstr = `SELECT dbo.NxDbGetRemnantsForNesting(@p1) AS OUT`
	var r0 string
	logf(tenantID, sqlstr, iPathID)
	if err := db.QueryRowContext(ctx, sqlstr, iPathID).Scan(&r0); err != nil {
		return "", logerror(err)
	}
	return r0, nil
}
