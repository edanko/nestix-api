package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxPPupdateSpdTimes calls the stored procedure 'dbo.NxPPupdateSpdTimes(int)' on db.
func NxPPupdateSpdTimes(ctx context.Context, db *sqlx.DB, iPathID int) error {
	// call dbo.NxPPupdateSpdTimes
	const sqlstr = `dbo.NxPPupdateSpdTimes`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iPathID", iPathID)); err != nil {
		return logerror(err)
	}
	return nil
}
