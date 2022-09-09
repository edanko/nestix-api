package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTupdateMatOrderline calls the stored procedure 'dbo.CUTupdateMatOrderline(int, float, float, float, float, int, int, int)' on db.
func CUTupdateMatOrderline(
	ctx context.Context,
	db *sqlx.DB,
	iSheetPathID int,
	lfWidth, lfLenght, lfThick, lfNestPct float64,
	iCount, iFreestockID, iCustomerID int,
) error {
	// call dbo.CUTupdateMatOrderline
	const sqlstr = `dbo.CUTupdateMatOrderline`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iSheetPathID", iSheetPathID), sql.Named("lfWidth", lfWidth), sql.Named("lfLenght", lfLenght), sql.Named("lfThick", lfThick), sql.Named("lfNestPct", lfNestPct), sql.Named("iCount", iCount), sql.Named("iFreestockID", iFreestockID), sql.Named("iCustomerID", iCustomerID)); err != nil {
		return logerror(err)
	}
	return nil
}
