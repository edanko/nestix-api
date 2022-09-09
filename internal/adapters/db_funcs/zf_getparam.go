package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// NxDbGetParam calls the stored function 'dbo.NxDbGetParam(nvarchar, int, int, int, int) nvarchar' on db.
func NxDbGetParam(
	ctx context.Context,
	db *sqlx.DB,
	strParamName string,
	iPhaseID, iWorkPlaceID, iWorkgroupID, iSiteID int,
) (string, error) {
	// call dbo.NxDbGetParam
	const sqlstr = `SELECT dbo.NxDbGetParam(@p1, @p2, @p3, @p4, @p5) AS OUT`
	var r0 string
	logf(tenantID, sqlstr, strParamName, iPhaseID, iWorkPlaceID, iWorkgroupID, iSiteID)
	if err := db.QueryRowContext(ctx, sqlstr, strParamName, iPhaseID, iWorkPlaceID, iWorkgroupID, iSiteID).Scan(&r0); err != nil {
		return "", logerror(err)
	}
	return r0, nil
}
