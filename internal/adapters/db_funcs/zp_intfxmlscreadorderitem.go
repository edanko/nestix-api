package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxIntfXMLSCReadOrderItem calls the stored procedure 'dbo.NxIntfXMLSCReadOrderItem(int, nvarchar, int, int, nvarchar)' on db.
func NxIntfXMLSCReadOrderItem(
	ctx context.Context,
	db *sqlx.DB,
	idoc int,
	strXMLPath string,
	iParentNodeID, iProductType int,
	strNewOrderNo string,
) error {
	// call dbo.NxIntfXMLSCReadOrderItem
	const sqlstr = `dbo.NxIntfXMLSCReadOrderItem`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("idoc", idoc), sql.Named("strXMLPath", strXMLPath), sql.Named("iParentNodeID", iParentNodeID), sql.Named("iProductType", iProductType), sql.Named("strNewOrderNo", strNewOrderNo)); err != nil {
		return logerror(err)
	}
	return nil
}
