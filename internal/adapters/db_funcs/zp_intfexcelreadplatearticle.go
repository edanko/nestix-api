package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxIntfExcelReadPlateArticle calls the stored procedure 'dbo.NxIntfExcelReadPlateArticle(int, nvarchar, int, nvarchar)' on db.
func NxIntfExcelReadPlateArticle(
	ctx context.Context,
	db *sqlx.DB,
	idoc int,
	strXMLPath string,
	iTransactNo int,
	strUserName string,
) error {
	// call dbo.NxIntfExcelReadPlateArticle
	const sqlstr = `dbo.NxIntfExcelReadPlateArticle`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("idoc", idoc), sql.Named("strXMLPath", strXMLPath), sql.Named("iTransactNo", iTransactNo), sql.Named("strUserName", strUserName)); err != nil {
		return logerror(err)
	}
	return nil
}
