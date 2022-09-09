package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxIntfXMLSCCreateSectionCopyMessage calls the stored procedure 'dbo.NxIntfXmlSCCreateSectionCopyMessage(nvarchar, nvarchar)' on db.
func NxIntfXMLSCCreateSectionCopyMessage(
	ctx context.Context,
	db *sqlx.DB,
	strOrderNo, strSection string,
) error {
	// call dbo.NxIntfXmlSCCreateSectionCopyMessage
	const sqlstr = `dbo.NxIntfXmlSCCreateSectionCopyMessage`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strOrderNo", strOrderNo), sql.Named("strSection", strSection)); err != nil {
		return logerror(err)
	}
	return nil
}
