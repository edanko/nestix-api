package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTintfXMLReadPlateItem calls the stored procedure 'dbo.CUTintfXMLReadPlateItem(int, nvarchar, nvarchar, nvarchar, nvarchar, int)' on db.
func CUTintfXMLReadPlateItem(
	ctx context.Context,
	db *sqlx.DB,
	idoc int,
	strXMLPath, strOrderNo, strParentName, strSection string,
	iNodeID int,
) error {
	// call dbo.CUTintfXMLReadPlateItem
	const sqlstr = `dbo.CUTintfXMLReadPlateItem`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("idoc", idoc), sql.Named("strXMLPath", strXMLPath), sql.Named("strOrderNo", strOrderNo), sql.Named("strParentName", strParentName), sql.Named("strSection", strSection), sql.Named("iNodeID", iNodeID)); err != nil {
		return logerror(err)
	}
	return nil
}
