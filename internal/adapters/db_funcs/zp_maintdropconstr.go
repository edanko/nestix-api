package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxMaintDropConstr calls the stored procedure 'dbo.nx_maint_drop_constr(nvarchar)' on db.
func NxMaintDropConstr(ctx context.Context, db *sqlx.DB, strTabname string) error {
	// call dbo.nx_maint_drop_constr
	const sqlstr = `dbo.nx_maint_drop_constr`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strTabname", strTabname)); err != nil {
		return logerror(err)
	}
	return nil
}
