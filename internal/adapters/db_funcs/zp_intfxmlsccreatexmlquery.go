package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxIntfXMLSCCreateXMLQuery calls the stored procedure 'dbo.NxIntfXmlSCCreateXMLQuery(nvarchar, nvarchar) nvarchar' on db.
func NxIntfXMLSCCreateXMLQuery(
	ctx context.Context,
	db *sqlx.DB,
	strElement, strSection string,
) (string, error) {
	// call dbo.NxIntfXmlSCCreateXMLQuery
	const sqlstr = `dbo.NxIntfXmlSCCreateXMLQuery`
	var strQuery string
	logf(tenantID, sqlstr, strElement, strSection)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strElement", strElement), sql.Named("strSection", strSection), sql.Named("strQuery", sql.Out{Dest: &strQuery})); err != nil {
		return "", logerror(err)
	}
	return strQuery, nil
}
