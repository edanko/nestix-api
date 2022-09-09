package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTupdatePosNestedCount calls the stored procedure 'dbo.CUTupdatePosNestedCount(int, nvarchar)' on db.
func CUTupdatePosNestedCount(ctx context.Context, db *sqlx.DB, iPosID int, strLang string) error {
	// call dbo.CUTupdatePosNestedCount
	const sqlstr = `dbo.CUTupdatePosNestedCount`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iPosID", iPosID), sql.Named("strLang", strLang)); err != nil {
		return logerror(err)
	}
	return nil
}
