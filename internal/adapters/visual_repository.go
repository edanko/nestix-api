package adapters

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/edanko/nestix-api/pkg/tenant"
)

// VisualModel represents a row from 'dbo.nxvisual'.
type VisualModel struct {
	ID            int64          `db:"nxvisualid"`
	OrderID       sql.NullInt64  `db:"nxorderlineid"`
	InventoryID   sql.NullInt64  `db:"nxinventoryid"`
	ProductID     sql.NullInt64  `db:"nxproductid"`
	PartID        sql.NullInt64  `db:"nxpartid"`
	PathID        sql.NullInt64  `db:"nxpathid"`
	User          sql.NullString `db:"nxvsuser"`
	ConnectionID  int64          `db:"nxconnectionid"`
	HostProcessID sql.NullInt64  `db:"nxhostprocessid"`
	HostName      sql.NullString `db:"nxhostname"`
	SiteName      sql.NullString `db:"nxsitename"`
}

type VisualRepository struct {
	dbs map[string]*sqlx.DB
}

func NewVisualRepository(db map[string]*sqlx.DB) *VisualRepository {
	return &VisualRepository{
		dbs: db,
	}
}

func (r *VisualRepository) List(ctx context.Context) ([]*VisualModel, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT
		nxvisualid, nxorderlineid, nxinventoryid, nxproductid, nxpartid, nxpathid, nxvsuser, nxconnectionid, nxhostprocessid, nxhostname, nxsitename
		FROM dbo.nxvisual`
	logf(tenantID, sqlstr)

	rows, err := r.dbs[tenantID].QueryxContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	var nn []*VisualModel
	for rows.Next() {
		var n VisualModel
		err := rows.StructScan(&n)
		if err != nil {
			return nil, logerror(err)
		}
		nn = append(nn, &n)
	}

	return nn, nil
}
