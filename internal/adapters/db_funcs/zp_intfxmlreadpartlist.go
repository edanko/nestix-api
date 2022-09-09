package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTintfXMLReadPartlist calls the stored procedure 'dbo.CUTintfXMLReadPartlist(int, nvarchar, int, nvarchar, nvarchar, nvarchar)' on db.
func CUTintfXMLReadPartlist(
	ctx context.Context,
	db *sqlx.DB,
	idoc int,
	strXMLPath string,
	iParentNodeID int,
	strOrderNo, strParentName, strSection string,
) error {
	// call dbo.CUTintfXMLReadPartlist
	const sqlstr = `dbo.CUTintfXMLReadPartlist`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("idoc", idoc), sql.Named("strXMLPath", strXMLPath), sql.Named("iParentNodeID", iParentNodeID), sql.Named("strOrderNo", strOrderNo), sql.Named("strParentName", strParentName), sql.Named("strSection", strSection)); err != nil {
		return logerror(err)
	}
	return nil
}
