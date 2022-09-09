package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxIntfExcelReadPlateOrderItem calls the stored procedure 'dbo.NxIntfExcelReadPlateOrderItem(int, nvarchar, int, nvarchar)' on db.
func NxIntfExcelReadPlateOrderItem(
	ctx context.Context,
	db *sqlx.DB,
	idoc int,
	strXMLPath string,
	iTransactNo int,
	strUserName string,
) error {
	// call dbo.NxIntfExcelReadPlateOrderItem
	const sqlstr = `dbo.NxIntfExcelReadPlateOrderItem`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("idoc", idoc), sql.Named("strXMLPath", strXMLPath), sql.Named("iTransactNo", iTransactNo), sql.Named("strUserName", strUserName)); err != nil {
		return logerror(err)
	}
	return nil
}
