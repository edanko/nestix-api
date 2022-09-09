package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTupdateNestToChanged calls the stored procedure 'dbo.CUTupdateNestToChanged(int, int)' on db.
func CUTupdateNestToChanged(ctx context.Context, db *sqlx.DB, iPartID, iChangeType int) error {
	// call dbo.CUTupdateNestToChanged
	const sqlstr = `dbo.CUTupdateNestToChanged`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iPartID", iPartID), sql.Named("iChangeType", iChangeType)); err != nil {
		return logerror(err)
	}
	return nil
}
