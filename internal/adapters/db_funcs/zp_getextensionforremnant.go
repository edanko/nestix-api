package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTgetExtensionForRemnant calls the stored procedure 'dbo.CUTgetExtensionForRemnant(int, int) nvarchar' on db.
func CUTgetExtensionForRemnant(
	ctx context.Context,
	db *sqlx.DB,
	iInventoryID, iSheetPathID int,
) (string, error) {
	// call dbo.CUTgetExtensionForRemnant
	const sqlstr = `dbo.CUTgetExtensionForRemnant`
	var strExtension string
	logf(tenantID, sqlstr, iInventoryID, iSheetPathID)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iInventoryID", iInventoryID), sql.Named("iSheetPathID", iSheetPathID), sql.Named("strExtension", sql.Out{Dest: &strExtension})); err != nil {
		return "", logerror(err)
	}
	return strExtension, nil
}
