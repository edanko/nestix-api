package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxIntfExcelProcessSQL calls the stored procedure 'dbo.NxIntfExcelProcessSQL(int, int, nvarchar, nvarchar, nvarchar)' on db.
func NxIntfExcelProcessSQL(
	ctx context.Context,
	db *sqlx.DB,
	iTransactNo, iTableProcessingWay int,
	strElementName, strTableName, strTableAlias string,
) error {
	// call dbo.NxIntfExcelProcessSQL
	const sqlstr = `dbo.NxIntfExcelProcessSQL`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iTransactNo", iTransactNo), sql.Named("iTableProcessingWay", iTableProcessingWay), sql.Named("strElementName", strElementName), sql.Named("strTableName", strTableName), sql.Named("strTableAlias", strTableAlias)); err != nil {
		return logerror(err)
	}
	return nil
}
