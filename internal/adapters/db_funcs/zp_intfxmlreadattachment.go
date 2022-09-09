package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTintfXMLReadAttachment calls the stored procedure 'dbo.CUTintfXMLReadAttachment(int, nvarchar, int, int)' on db.
func CUTintfXMLReadAttachment(
	ctx context.Context,
	db *sqlx.DB,
	idoc int,
	strXMLPath string,
	iOrderlineID, iParentNodeID int,
) error {
	// call dbo.CUTintfXMLReadAttachment
	const sqlstr = `dbo.CUTintfXMLReadAttachment`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("idoc", idoc), sql.Named("strXMLPath", strXMLPath), sql.Named("iOrderlineID", iOrderlineID), sql.Named("iParentNodeID", iParentNodeID)); err != nil {
		return logerror(err)
	}
	return nil
}
