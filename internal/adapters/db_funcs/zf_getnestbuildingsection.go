package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// NxDbGetNestBuildingSection calls the stored function 'dbo.NxDbGetNestBuildingSection(int) nvarchar' on db.
func NxDbGetNestBuildingSection(ctx context.Context, db *sqlx.DB, iPathID int) (string, error) {
	// call dbo.NxDbGetNestBuildingSection
	const sqlstr = `SELECT dbo.NxDbGetNestBuildingSection(@p1) AS OUT`
	var r0 string
	logf(tenantID, sqlstr, iPathID)
	if err := db.QueryRowContext(ctx, sqlstr, iPathID).Scan(&r0); err != nil {
		return "", logerror(err)
	}
	return r0, nil
}
