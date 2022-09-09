package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// BVLgetBevelCorrData calls the stored procedure 'dbo.BVLgetBevelCorrData(nvarchar, nvarchar, float, nchar, float, float, float)' on db.
func BVLgetBevelCorrData(
	ctx context.Context,
	db *sqlx.DB,
	strTech, strMatgroup string,
	lfThick float64,
	strBevelType string,
	lfBevelAngle, lfRootWidth, lfEdgeDist float64,
) error {
	// call dbo.BVLgetBevelCorrData
	const sqlstr = `dbo.BVLgetBevelCorrData`
	logf(tenantID, sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strTech", strTech), sql.Named("strMatgroup", strMatgroup), sql.Named("lfThick", lfThick), sql.Named("strBevelType", strBevelType), sql.Named("lfBevelAngle", lfBevelAngle), sql.Named("lfRootWidth", lfRootWidth), sql.Named("lfEdgeDist", lfEdgeDist)); err != nil {
		return logerror(err)
	}
	return nil
}
