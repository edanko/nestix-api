package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxDbGroupItemsToString calls the stored procedure 'dbo.NxDbGroupItemsToString(nvarchar, nvarchar, nvarchar, nvarchar, nvarchar, int)' on db.
func NxDbGroupItemsToString(
	ctx context.Context,
	db *sqlx.DB,
	strColumnToGroupBy, strColumnToList, strFROM, strWHERE, strSeparator string,
	iOrderingType int,
) error {
	// call dbo.NxDbGroupItemsToString
	const sqlstr = `dbo.NxDbGroupItemsToString`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strColumnToGroupBy", strColumnToGroupBy), sql.Named("strColumnToList", strColumnToList), sql.Named("strFROM", strFROM), sql.Named("strWHERE", strWHERE), sql.Named("strSeparator", strSeparator), sql.Named("iOrderingType", iOrderingType)); err != nil {
		return logerror(err)
	}
	return nil
}
