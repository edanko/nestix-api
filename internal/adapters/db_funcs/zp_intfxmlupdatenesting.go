package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTintfXMLUpdateNesting calls the stored procedure 'dbo.CUTintfXMLUpdateNesting(int, int)' on db.
func CUTintfXMLUpdateNesting(ctx context.Context, db *sqlx.DB, iOrderlineID, iUpdateType int) error {
	// call dbo.CUTintfXMLUpdateNesting
	const sqlstr = `dbo.CUTintfXMLUpdateNesting`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iOrderlineID", iOrderlineID), sql.Named("iUpdateType", iUpdateType)); err != nil {
		return logerror(err)
	}
	return nil
}
