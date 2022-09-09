package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxMaintGetObjectsContaining calls the stored procedure 'dbo.nx_maint_getObjectsContaining(sysname, int) int' on db.
func NxMaintGetObjectsContaining(
	ctx context.Context,
	db *sqlx.DB,
	fldname string,
	iMode int,
) (int, error) {
	// call dbo.nx_maint_getObjectsContaining
	const sqlstr = `dbo.nx_maint_getObjectsContaining`
	var fldcount int
	logf(tenantID, sqlstr, fldname, iMode)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("fldname", fldname), sql.Named("iMode", iMode), sql.Named("fldcount", sql.Out{Dest: &fldcount})); err != nil {
		return 0, logerror(err)
	}
	return fldcount, nil
}
