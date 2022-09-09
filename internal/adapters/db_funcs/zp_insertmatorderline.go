package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTinsertMatOrderline calls the stored procedure 'dbo.CUTinsertMatOrderline(nvarchar, float, float, float, float, int, int, int)' on db.
func CUTinsertMatOrderline(
	ctx context.Context,
	db *sqlx.DB,
	strUser string,
	lfWidth, lfLenght, lfThick, lfNestPct float64,
	iOrderCnt, iProductID, iInventoryID int,
) error {
	// call dbo.CUTinsertMatOrderline
	const sqlstr = `dbo.CUTinsertMatOrderline`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strUser", strUser), sql.Named("lfWidth", lfWidth), sql.Named("lfLenght", lfLenght), sql.Named("lfThick", lfThick), sql.Named("lfNestPct", lfNestPct), sql.Named("iOrderCnt", iOrderCnt), sql.Named("iProductID", iProductID), sql.Named("iInventoryID", iInventoryID)); err != nil {
		return logerror(err)
	}
	return nil
}
