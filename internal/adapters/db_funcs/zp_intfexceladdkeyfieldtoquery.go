package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxIntfExcelAddKeyFieldToQuery calls the stored procedure 'dbo.NxIntfExcelAddKeyFieldToQuery(nvarchar, nvarchar, nvarchar, nvarchar, nvarchar, int, nvarchar)' on db.
func NxIntfExcelAddKeyFieldToQuery(
	ctx context.Context,
	db *sqlx.DB,
	strDestinationColumn, strSourceAlias, strElementName, strTableName, strTableAlias string,
	iTransactNo int,
	strFieldValue string,
) error {
	// call dbo.NxIntfExcelAddKeyFieldToQuery
	const sqlstr = `dbo.NxIntfExcelAddKeyFieldToQuery`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strDestinationColumn", strDestinationColumn), sql.Named("strSourceAlias", strSourceAlias), sql.Named("strElementName", strElementName), sql.Named("strTableName", strTableName), sql.Named("strTableAlias", strTableAlias), sql.Named("iTransactNo", iTransactNo), sql.Named("strFieldValue", strFieldValue)); err != nil {
		return logerror(err)
	}
	return nil
}
