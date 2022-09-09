package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTdeletePathMaterialOrderlines calls the stored procedure 'dbo.CUTdeletePathMaterialOrderlines(int, nvarchar)' on db.
func CUTdeletePathMaterialOrderlines(
	ctx context.Context,
	db *sqlx.DB,
	iPathID int,
	strLang string,
) error {
	// call dbo.CUTdeletePathMaterialOrderlines
	const sqlstr = `dbo.CUTdeletePathMaterialOrderlines`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("iPathID", iPathID), sql.Named("strLang", strLang)); err != nil {
		return logerror(err)
	}
	return nil
}
