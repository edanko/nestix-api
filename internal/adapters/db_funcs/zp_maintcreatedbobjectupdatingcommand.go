package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxMaintCreateDbObjectUpdatingCommand calls the stored procedure 'dbo.nx_maint_create_db_object_updating_command(nvarchar, int, int)' on db.
func NxMaintCreateDbObjectUpdatingCommand(
	ctx context.Context,
	db *sqlx.DB,
	stdData string,
	iUsesUpdateChecked, iCreateTransactionCmd int,
) error {
	// call dbo.nx_maint_create_db_object_updating_command
	const sqlstr = `dbo.nx_maint_create_db_object_updating_command`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("stdData", stdData), sql.Named("iUsesUpdateChecked", iUsesUpdateChecked), sql.Named("iCreateTransactionCmd", iCreateTransactionCmd)); err != nil {
		return logerror(err)
	}
	return nil
}
