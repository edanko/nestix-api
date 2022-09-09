package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxIntfXMLSCReadMaterialOrder calls the stored procedure 'dbo.NxIntfXMLSCReadMaterialOrder(int, nvarchar, nvarchar)' on db.
func NxIntfXMLSCReadMaterialOrder(
	ctx context.Context,
	db *sqlx.DB,
	idoc int,
	strXMLPath, strNewOrderNo string,
) error {
	// call dbo.NxIntfXMLSCReadMaterialOrder
	const sqlstr = `dbo.NxIntfXMLSCReadMaterialOrder`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("idoc", idoc), sql.Named("strXMLPath", strXMLPath), sql.Named("strNewOrderNo", strNewOrderNo)); err != nil {
		return logerror(err)
	}
	return nil
}
