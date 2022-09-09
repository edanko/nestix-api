package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTintfXMLCreateNestingMessage calls the stored procedure 'dbo.CUTintfXmlCreateNestingMessage(nvarchar, int)' on db.
func CUTintfXMLCreateNestingMessage(
	ctx context.Context,
	db *sqlx.DB,
	strMsgSeqNo string,
	iPathID int,
) error {
	// call dbo.CUTintfXmlCreateNestingMessage
	const sqlstr = `dbo.CUTintfXmlCreateNestingMessage`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strMsgSeqNo", strMsgSeqNo), sql.Named("iPathID", iPathID)); err != nil {
		return logerror(err)
	}
	return nil
}
