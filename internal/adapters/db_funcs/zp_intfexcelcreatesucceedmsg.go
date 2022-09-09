package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxIntfExcelCreateSucceedMsg calls the stored procedure 'dbo.NxIntfExcelCreateSucceedMsg(int) nvarchar' on db.
func NxIntfExcelCreateSucceedMsg(ctx context.Context, db *sqlx.DB, iTransactNo int) (string, error) {
	// call dbo.NxIntfExcelCreateSucceedMsg
	const sqlstr = `dbo.NxIntfExcelCreateSucceedMsg`
	var strErrorMsg string
	logf(tenantID, sqlstr, iTransactNo)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iTransactNo", iTransactNo), sql.Named("strErrorMsg", sql.Out{Dest: &strErrorMsg})); err != nil {
		return "", logerror(err)
	}
	return strErrorMsg, nil
}
