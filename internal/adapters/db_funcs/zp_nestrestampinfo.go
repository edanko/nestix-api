package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTnestRestampInfo calls the stored procedure 'dbo.CUTnestRestampInfo(nvarchar, int, int, int, nvarchar)' on db.
func CUTnestRestampInfo(
	ctx context.Context,
	db *sqlx.DB,
	strKeyword string,
	iPathID, iPosID, iPartID int,
	strLang string,
) error {
	// call dbo.CUTnestRestampInfo
	const sqlstr = `dbo.CUTnestRestampInfo`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strKeyword", strKeyword), sql.Named("iPathID", iPathID), sql.Named("iPosID", iPosID), sql.Named("iPartID", iPartID), sql.Named("strLang", strLang)); err != nil {
		return logerror(err)
	}
	return nil
}
