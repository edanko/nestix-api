package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxMaintCreateDataUpdateScript calls the stored procedure 'dbo.nx_maint_CreateDataUpdateScript(nvarchar, nvarchar, nvarchar, nvarchar, int, int)' on db.
func NxMaintCreateDataUpdateScript(
	ctx context.Context,
	db *sqlx.DB,
	strTableSchema, strTableName, strCondition, strUniqueFields string,
	iUpdateType, iCreateTransactionCmd int,
) error {
	// call dbo.nx_maint_CreateDataUpdateScript
	const sqlstr = `dbo.nx_maint_CreateDataUpdateScript`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strTableSchema", strTableSchema), sql.Named("strTableName", strTableName), sql.Named("strCondition", strCondition), sql.Named("strUniqueFields", strUniqueFields), sql.Named("iUpdateType", iUpdateType), sql.Named("iCreateTransactionCmd", iCreateTransactionCmd)); err != nil {
		return logerror(err)
	}
	return nil
}
