package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxIntfXMLSCMatOrdersOfCustOrder calls the stored procedure 'dbo.NxIntfXMLSCMatOrdersOfCustOrder(nvarchar, nvarchar, int)' on db.
func NxIntfXMLSCMatOrdersOfCustOrder(
	ctx context.Context,
	db *sqlx.DB,
	strOrderNo, strSection string,
	iUpdTmpTable int,
) error {
	// call dbo.NxIntfXMLSCMatOrdersOfCustOrder
	const sqlstr = `dbo.NxIntfXMLSCMatOrdersOfCustOrder`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strOrderNo", strOrderNo), sql.Named("strSection", strSection), sql.Named("iUpdTmpTable", iUpdTmpTable)); err != nil {
		return logerror(err)
	}
	return nil
}
