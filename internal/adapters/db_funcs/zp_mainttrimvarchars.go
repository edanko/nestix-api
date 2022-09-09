package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// NxMaintTrimVarchars calls the stored procedure 'dbo.nx_maint_trim_varchars()' on db.
func NxMaintTrimVarchars(ctx context.Context, db *sqlx.DB) error {
	// call dbo.nx_maint_trim_varchars
	const sqlstr = `dbo.nx_maint_trim_varchars`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr); err != nil {
		return logerror(err)
	}
	return nil
}
