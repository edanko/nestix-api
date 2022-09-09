package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxIntfXMLSCReadProduct calls the stored procedure 'dbo.NxIntfXMLSCReadProduct(int, varchar, int) int' on db.
func NxIntfXMLSCReadProduct(
	ctx context.Context,
	db *sqlx.DB,
	idoc int,
	strXMLPath string,
	iParentNodeID int,
) (int, error) {
	// call dbo.NxIntfXMLSCReadProduct
	const sqlstr = `dbo.NxIntfXMLSCReadProduct`
	var iProductID int
	logf(tenantID, sqlstr, idoc, strXMLPath, iParentNodeID)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("idoc", idoc), sql.Named("strXMLPath", strXMLPath), sql.Named("iParentNodeID", iParentNodeID), sql.Named("iProductID", sql.Out{Dest: &iProductID})); err != nil {
		return 0, logerror(err)
	}
	return iProductID, nil
}
