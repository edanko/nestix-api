package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// NxDbLangTranslateProc calls the stored procedure 'dbo.NxDbLangTranslateProc(nvarchar, nvarchar) nvarchar' on db.
func NxDbLangTranslateProc(ctx context.Context, db *sqlx.DB, strOrgText, strLang string) (string, error) {
	// call dbo.NxDbLangTranslateProc
	const sqlstr = `dbo.NxDbLangTranslateProc`
	var strNewText string
	logf(tenantID, sqlstr, strOrgText, strLang)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("strOrgText", strOrgText), sql.Named("strLang", strLang), sql.Named("strNewText", sql.Out{Dest: &strNewText})); err != nil {
		return "", logerror(err)
	}
	return strNewText, nil
}
