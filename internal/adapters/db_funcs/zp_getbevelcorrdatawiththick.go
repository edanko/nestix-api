package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// BVLgetBevelCorrDataWithThick calls the stored procedure 'dbo.BVLgetBevelCorrDataWithThick(nvarchar, nvarchar, float, nchar, float, float, float, int) (float, float, float, float)' on db.
func BVLgetBevelCorrDataWithThick(
	ctx context.Context,
	db *sqlx.DB,
	strTech, strMatgroup string,
	lfThick float64,
	strBevelType string,
	lfBevelAngle, lfRootWidth, lfEdgeDist float64,
	iGetWithLTEAngle int,
) (float64, float64, float64, float64, error) {
	// call dbo.BVLgetBevelCorrDataWithThick
	const sqlstr = `dbo.BVLgetBevelCorrDataWithThick`
	var lfAngle float64
	var lfCorrAngle float64
	var lfCorrWidth float64
	var lfSpeed float64
	logf(tenantID,
		sqlstr,
		strTech,
		strMatgroup,
		lfThick,
		strBevelType,
		lfBevelAngle,
		lfRootWidth,
		lfEdgeDist,
		iGetWithLTEAngle,
	)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strTech", strTech), sql.Named("strMatgroup", strMatgroup), sql.Named("lfThick", lfThick), sql.Named("strBevelType", strBevelType), sql.Named("lfBevelAngle", lfBevelAngle), sql.Named("lfRootWidth", lfRootWidth), sql.Named("lfEdgeDist", lfEdgeDist), sql.Named("iGetWithLTEAngle", iGetWithLTEAngle), sql.Named("lfAngle", sql.Out{Dest: &lfAngle}), sql.Named("lfCorrAngle", sql.Out{Dest: &lfCorrAngle}), sql.Named("lfCorrWidth", sql.Out{Dest: &lfCorrWidth}), sql.Named("lfSpeed", sql.Out{Dest: &lfSpeed})); err != nil {
		return 0.0, 0.0, 0.0, 0.0, logerror(err)
	}
	return lfAngle, lfCorrAngle, lfCorrWidth, lfSpeed, nil
}
