package db_funcs

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CUTintfXMLReadInnerXML calls the stored procedure 'dbo.CUTintfXMLReadInnerXml(int, nvarchar, nvarchar, int) xml' on db.
func CUTintfXMLReadInnerXML(
	ctx context.Context,
	db *sqlx.DB,
	idoc int,
	strXMLPath, strXMLElementName string,
	iNodeID int,
) ([]byte, error) {
	// call dbo.CUTintfXMLReadInnerXml
	const sqlstr = `dbo.CUTintfXMLReadInnerXml`
	var xmlInnerXML []byte
	logf(tenantID, sqlstr, idoc, strXMLPath, strXMLElementName, iNodeID)
	if _, err := db.ExecContext(ctx, sqlstr, sql.Named("idoc", idoc), sql.Named("strXMLPath", strXMLPath), sql.Named("strXMLElementName", strXMLElementName), sql.Named("iNodeID", iNodeID), sql.Named("xmlInnerXml", sql.Out{Dest: &xmlInnerXML})); err != nil {
		return nil, logerror(err)
	}
	return xmlInnerXML, nil
}
