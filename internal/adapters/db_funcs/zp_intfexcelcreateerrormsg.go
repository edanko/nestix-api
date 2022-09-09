package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxIntfExcelCreateErrorMsg calls the stored procedure 'dbo.NxIntfExcelCreateErrorMsg(int)' on db.
func NxIntfExcelCreateErrorMsg(ctx context.Context, db *sqlx.DB, iTransactNo int) error {
	// call dbo.NxIntfExcelCreateErrorMsg
	const sqlstr = `dbo.NxIntfExcelCreateErrorMsg`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iTransactNo", iTransactNo)); err != nil {
		return logerror(err)
	}
	return nil
}
