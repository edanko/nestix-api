package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxIntfExcelCreateQuery calls the stored procedure 'dbo.NxIntfExcelCreateQuery(int, nvarchar, nvarchar, nvarchar, nvarchar, nvarchar, int)' on db.
func NxIntfExcelCreateQuery(
	ctx context.Context,
	db *sqlx.DB,
	idoc int,
	strXMLPath, strIntfType, strUniqueFields, strTableName, strTableAliasName string,
	iTransactNo int,
) error {
	// call dbo.NxIntfExcelCreateQuery
	const sqlstr = `dbo.NxIntfExcelCreateQuery`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("idoc", idoc), sql.Named("strXMLPath", strXMLPath), sql.Named("strIntfType", strIntfType), sql.Named("strUniqueFields", strUniqueFields), sql.Named("strTableName", strTableName), sql.Named("strTableAliasName", strTableAliasName), sql.Named("iTransactNo", iTransactNo)); err != nil {
		return logerror(err)
	}
	return nil
}
