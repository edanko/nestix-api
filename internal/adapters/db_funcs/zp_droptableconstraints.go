package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxDropTableConstraints calls the stored procedure 'dbo.nx_drop_table_constraints(nvarchar)' on db.
func NxDropTableConstraints(ctx context.Context, db *sqlx.DB, strTabname string) error {
	// call dbo.nx_drop_table_constraints
	const sqlstr = `dbo.nx_drop_table_constraints`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strTabname", strTabname)); err != nil {
		return logerror(err)
	}
	return nil
}
