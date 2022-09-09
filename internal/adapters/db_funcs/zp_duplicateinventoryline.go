package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxDbDuplicateInventoryLine calls the stored procedure 'dbo.NxDbDuplicateInventoryLine(int, nvarchar) int' on db.
func NxDbDuplicateInventoryLine(
	ctx context.Context,
	db *sqlx.DB,
	iInventoryID int,
	strExcludeColumns string,
) (int, error) {
	// call dbo.NxDbDuplicateInventoryLine
	const sqlstr = `dbo.NxDbDuplicateInventoryLine`
	var iNewInventoryID int
	logf(tenantID, sqlstr, iInventoryID, strExcludeColumns)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iInventoryID", iInventoryID), sql.Named("strExcludeColumns", strExcludeColumns), sql.Named("iNewInventoryID", sql.Out{Dest: &iNewInventoryID})); err != nil {
		return 0, logerror(err)
	}
	return iNewInventoryID, nil
}
