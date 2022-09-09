package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTintfXMLReadDrawing calls the stored procedure 'dbo.CUTintfXMLReadDrawing(int, nvarchar, int) int' on db.
func CUTintfXMLReadDrawing(
	ctx context.Context,
	db *sqlx.DB,
	idoc int,
	strXMLPath string,
	iNodeID int,
) (int, error) {
	// call dbo.CUTintfXMLReadDrawing
	const sqlstr = `dbo.CUTintfXMLReadDrawing`
	var iPartID int
	logf(tenantID, sqlstr, idoc, strXMLPath, iNodeID)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("idoc", idoc), sql.Named("strXMLPath", strXMLPath), sql.Named("iNodeID", iNodeID), sql.Named("iPartID", sql.Out{Dest: &iPartID})); err != nil {
		return 0, logerror(err)
	}
	return iPartID, nil
}
