package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxMaintCreateDataCopyScript calls the stored procedure 'dbo.nx_maint_CreateDataCopyScript(nvarchar, nvarchar)' on db.
func NxMaintCreateDataCopyScript(
	ctx context.Context,
	db *sqlx.DB,
	strTableName, strCondition string,
) error {
	// call dbo.nx_maint_CreateDataCopyScript
	const sqlstr = `dbo.nx_maint_CreateDataCopyScript`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strTableName", strTableName), sql.Named("strCondition", strCondition)); err != nil {
		return logerror(err)
	}
	return nil
}
