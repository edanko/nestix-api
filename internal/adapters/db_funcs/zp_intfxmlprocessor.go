package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTintfXMLProcessor calls the stored procedure 'dbo.CUTintfXMLProcessor(nvarchar, xml)' on db.
func CUTintfXMLProcessor(
	ctx context.Context,
	db *sqlx.DB,
	strFilename string,
	xmlInputData []byte,
) error {
	// call dbo.CUTintfXMLProcessor
	const sqlstr = `dbo.CUTintfXMLProcessor`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strFilename", strFilename), sql.Named("xmlInputData", xmlInputData)); err != nil {
		return logerror(err)
	}
	return nil
}
