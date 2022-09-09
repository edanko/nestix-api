package db_funcs

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// NxDbLangTranslateFunc calls the stored function 'dbo.NxDbLangTranslateFunc(nvarchar, nvarchar) nvarchar' on db.
func NxDbLangTranslateFunc(ctx context.Context, db *sqlx.DB, strOrgText, strLang string) (string, error) {
	// call dbo.NxDbLangTranslateFunc
	const sqlstr = `SELECT dbo.NxDbLangTranslateFunc(@p1, @p2) AS OUT`
	var r0 string
	logf(tenantID, sqlstr, strOrgText, strLang)
	if err := db.QueryRowContext(ctx, sqlstr, strOrgText, strLang).Scan(&r0); err != nil {
		return "", logerror(err)
	}
	return r0, nil
}
