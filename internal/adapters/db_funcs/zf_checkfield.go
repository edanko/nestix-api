package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// NxDbCheckField calls the stored function 'dbo.NxDbCheckField(nvarchar, nvarchar) nvarchar' on db.
func NxDbCheckField(
	ctx context.Context,
	db *sqlx.DB,
	strFieldName, strFieldValue string,
) (string, error) {
	// call dbo.NxDbCheckField
	const sqlstr = `SELECT dbo.NxDbCheckField(@p1, @p2) AS OUT`
	var r0 string
	logf(tenantID, sqlstr, strFieldName, strFieldValue)
	if err := db.QueryRowContext(ctx, sqlstr, strFieldName, strFieldValue).Scan(&r0); err != nil {
		return "", logerror(err)
	}
	return r0, nil
}
