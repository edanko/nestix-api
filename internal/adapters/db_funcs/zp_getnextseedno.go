package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxDbGetNextSeedNo calls the stored procedure 'dbo.NxDbGetNextSeedNo(nvarchar) nvarchar' on db.
func NxDbGetNextSeedNo(ctx context.Context, db *sqlx.DB, strSeedName string) (string, error) {
	// call dbo.NxDbGetNextSeedNo
	const sqlstr = `dbo.NxDbGetNextSeedNo`
	var strOutput string
	logf(tenantID, sqlstr, strSeedName)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strSeedName", strSeedName), sql.Named("strOutput", sql.Out{Dest: &strOutput})); err != nil {
		return "", logerror(err)
	}
	return strOutput, nil
}
