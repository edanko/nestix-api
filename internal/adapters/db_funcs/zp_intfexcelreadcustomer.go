package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxIntfExcelReadCustomer calls the stored procedure 'dbo.NxIntfExcelReadCustomer(int, nvarchar, int, nvarchar)' on db.
func NxIntfExcelReadCustomer(
	ctx context.Context,
	db *sqlx.DB,
	idoc int,
	strXMLPath string,
	iTransactNo int,
	strUserName string,
) error {
	// call dbo.NxIntfExcelReadCustomer
	const sqlstr = `dbo.NxIntfExcelReadCustomer`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("idoc", idoc), sql.Named("strXMLPath", strXMLPath), sql.Named("iTransactNo", iTransactNo), sql.Named("strUserName", strUserName)); err != nil {
		return logerror(err)
	}
	return nil
}
