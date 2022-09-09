package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxIntfXMLSCReadSectionCopy calls the stored procedure 'dbo.NxIntfXMLSCReadSectionCopy(nvarchar, xml, nvarchar, int, int, int)' on db.
func NxIntfXMLSCReadSectionCopy(
	ctx context.Context,
	db *sqlx.DB,
	strNewOrderNo string,
	xmlInputData []byte,
	strSection string,
	iIsProductRead, iIsUnNestedPartRead, iIsUnDrawnPartRead int,
) error {
	// call dbo.NxIntfXMLSCReadSectionCopy
	const sqlstr = `dbo.NxIntfXMLSCReadSectionCopy`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strNewOrderNo", strNewOrderNo), sql.Named("xmlInputData", xmlInputData), sql.Named("strSection", strSection), sql.Named("iIsProductRead", iIsProductRead), sql.Named("iIsUnNestedPartRead", iIsUnNestedPartRead), sql.Named("iIsUnDrawnPartRead", iIsUnDrawnPartRead)); err != nil {
		return logerror(err)
	}
	return nil
}
