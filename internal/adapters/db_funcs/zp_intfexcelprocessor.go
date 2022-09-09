package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxIntfExcelProcessor calls the stored procedure 'dbo.NxIntfExcelProcessor(nvarchar, xml, int)' on db.
func NxIntfExcelProcessor(
	ctx context.Context,
	db *sqlx.DB,
	strFilename string,
	xmlInputData []byte,
	iIsDatabaseCheck int,
) error {
	// call dbo.NxIntfExcelProcessor
	const sqlstr = `dbo.NxIntfExcelProcessor`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strFilename", strFilename), sql.Named("xmlInputData", xmlInputData), sql.Named("iIsDatabaseCheck", iIsDatabaseCheck)); err != nil {
		return logerror(err)
	}
	return nil
}
