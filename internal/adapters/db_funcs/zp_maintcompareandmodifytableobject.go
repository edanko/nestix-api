package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxMaintCompareAndModifyTableObject calls the stored procedure 'dbo.nx_maint_compare_and_modify_table_object(nvarchar, nvarchar, nvarchar, nvarchar, nvarchar, nvarchar, nvarchar, nvarchar, int)' on db.
func NxMaintCompareAndModifyTableObject(
	ctx context.Context,
	db *sqlx.DB,
	strObjectType, strObjectSchema, strObjectTable, strObjectColumn, strObjectDataType, strObjectDataLength, strObjectIsNullable, strDefaultValue string,
	iUpdateDb int,
) error {
	// call dbo.nx_maint_compare_and_modify_table_object
	const sqlstr = `dbo.nx_maint_compare_and_modify_table_object`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strObjectType", strObjectType), sql.Named("strObjectSchema", strObjectSchema), sql.Named("strObjectTable", strObjectTable), sql.Named("strObjectColumn", strObjectColumn), sql.Named("strObjectDataType", strObjectDataType), sql.Named("strObjectDataLength", strObjectDataLength), sql.Named("strObjectIsNullable", strObjectIsNullable), sql.Named("strDefaultValue", strDefaultValue), sql.Named("iUpdateDb", iUpdateDb)); err != nil {
		return logerror(err)
	}
	return nil
}
