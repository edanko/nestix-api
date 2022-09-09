package adapters

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/edanko/nestix-api/internal/domain/sheetpath"
	"github.com/edanko/nestix-api/pkg/tenant"
)

// SheetPathModel represents a row from 'dbo.nxsheetpath'.
type SheetPathModel struct {
	ID                int64           `db:"nxsheetpathid"`
	PathID            sql.NullInt64   `db:"nxpathid"`
	MatOrderID        int64           `db:"nxmatorderlineid"`
	InventoryID       sql.NullInt64   `db:"nxinventoryid"`
	NestReservationID sql.NullInt64   `db:"nxnestreservationid"`
	ProductID         sql.NullInt64   `db:"nxproductid"`
	PartID            sql.NullInt64   `db:"nxpartid"`
	CutLength         sql.NullFloat64 `db:"nxcutlength"`
	PowderLength      sql.NullFloat64 `db:"nxpowderlength"`
	RapidLength       sql.NullFloat64 `db:"nxrapidlength"`
	NetArea           sql.NullFloat64 `db:"nxspnetarea"`
	UsedArea          sql.NullFloat64 `db:"nxusedarea"`
	Used              sql.NullFloat64 `db:"nxused"`
	MarkCount         sql.NullInt64   `db:"nxmarkcount"`
	ProfCount         sql.NullInt64   `db:"nxprofcount"`
	IsChanged         sql.NullInt64   `db:"nxspischanged"`
	InsertDate        sql.NullTime    `db:"nxspinsertdate"`
}

type SheetPathRepository struct {
	dbs map[string]*sqlx.DB
}

func NewSheetPathRepository(db map[string]*sqlx.DB) *SheetPathRepository {
	return &SheetPathRepository{
		dbs: db,
	}
}

// Insert inserts the SheetPathModel to the database.
func (r *SheetPathRepository) Insert(ctx context.Context, n *SheetPathModel) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `INSERT INTO dbo.nxsheetpath (
		nxsheetpathid, nxpathid, nxmatorderlineid, nxinventoryid, nxnestreservationid, nxproductid, nxpartid, nxcutlength, nxpowderlength, nxrapidlength, nxspnetarea, nxusedarea, nxused, nxmarkcount, nxprofcount, nxspischanged, nxspinsertdate
		) VALUES (
		@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, @p13, @p14, @p15, @p16, @p17
		)`
	logf(tenantID,
		sqlstr,
		n.ID,
		n.PathID,
		n.MatOrderID,
		n.InventoryID,
		n.NestReservationID,
		n.ProductID,
		n.PartID,
		n.CutLength,
		n.PowderLength,
		n.RapidLength,
		n.NetArea,
		n.UsedArea,
		n.Used,
		n.MarkCount,
		n.ProfCount,
		n.IsChanged,
		n.InsertDate,
	)
	if _, err := r.dbs[tenantID].ExecContext(
		ctx,
		sqlstr,
		n.ID,
		n.PathID,
		n.MatOrderID,
		n.InventoryID,
		n.NestReservationID,
		n.ProductID,
		n.PartID,
		n.CutLength,
		n.PowderLength,
		n.RapidLength,
		n.NetArea,
		n.UsedArea,
		n.Used,
		n.MarkCount,
		n.ProfCount,
		n.IsChanged,
		n.InsertDate,
	); err != nil {
		return logerror(err)
	}
	return nil
}

// Update updates a SheetPathModel in the database.
func (r *SheetPathRepository) Update(ctx context.Context, n *SheetPathModel) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `UPDATE dbo.nxsheetpath SET ` +
		`nxpathid = @p1, nxmatorderlineid = @p2, nxinventoryid = @p3, nxnestreservationid = @p4, nxproductid = @p5, nxpartid = @p6, nxcutlength = @p7, nxpowderlength = @p8, nxrapidlength = @p9, nxspnetarea = @p10, nxusedarea = @p11, nxused = @p12, nxmarkcount = @p13, nxprofcount = @p14, nxspischanged = @p15, nxspinsertdate = @p16 ` +
		`WHERE nxsheetpathid = @p17`
	logf(
		tenantID,
		sqlstr,
		n.PathID,
		n.MatOrderID,
		n.InventoryID,
		n.NestReservationID,
		n.ProductID,
		n.PartID,
		n.CutLength,
		n.PowderLength,
		n.RapidLength,
		n.NetArea,
		n.UsedArea,
		n.Used,
		n.MarkCount,
		n.ProfCount,
		n.IsChanged,
		n.InsertDate,
		n.ID,
	)
	if _, err := r.dbs[tenantID].ExecContext(
		ctx,
		sqlstr,
		n.PathID,
		n.MatOrderID,
		n.InventoryID,
		n.NestReservationID,
		n.ProductID,
		n.PartID,
		n.CutLength,
		n.PowderLength,
		n.RapidLength,
		n.NetArea,
		n.UsedArea,
		n.Used,
		n.MarkCount,
		n.ProfCount,
		n.IsChanged,
		n.InsertDate,
		n.ID,
	); err != nil {
		return logerror(err)
	}
	return nil
}

// Delete deletes the SheetPathModel from the database.
func (r *SheetPathRepository) DeleteByID(ctx context.Context, id int64) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `DELETE FROM dbo.nxsheetpath
		WHERE nxsheetpathid = @p1`
	logf(tenantID, sqlstr, id)
	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, id); err != nil {
		return logerror(err)
	}
	return nil
}

// Delete deletes the SheetPathModel from the database.
func (r *SheetPathRepository) DeleteByIDs(ctx context.Context, ids []int64) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `DELETE FROM dbo.nxsheetpath
		WHERE nxsheetpathid IN (?)`
	query, args, err := sqlx.In(sqlstr, ids)
	if err != nil {
		return logerror(err)
	}
	query = r.dbs[tenantID].Rebind(query)

	logf(tenantID, query, ids)

	if _, err := r.dbs[tenantID].ExecContext(ctx, query, args...); err != nil {
		return logerror(err)
	}
	return nil
}

// Get retrieves a row from 'dbo.nxsheetpath' as a SheetPathModel.
func (r *SheetPathRepository) GetByID(ctx context.Context, id int64) (*sheetpath.SheetPath, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT
		nxsheetpathid, nxpathid, nxmatorderlineid, nxinventoryid, nxnestreservationid, nxproductid, nxpartid, nxcutlength, nxpowderlength, nxrapidlength, nxspnetarea, nxusedarea, nxused, nxmarkcount, nxprofcount, nxspischanged, nxspinsertdate
		FROM dbo.nxsheetpath
		WHERE nxsheetpathid = @p1`
	logf(tenantID, sqlstr, id)
	n := SheetPathModel{}
	row := r.dbs[tenantID].QueryRowxContext(ctx, sqlstr, id)
	err := row.StructScan(&n)
	if err != nil {
		return nil, logerror(err)
	}
	return mapSheetPathModel(&n), nil
}

func (r *SheetPathRepository) GetByPathID(ctx context.Context, id int64) (*sheetpath.SheetPath, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT
		nxsheetpathid, nxpathid, nxmatorderlineid, nxinventoryid, nxnestreservationid, nxproductid, nxpartid, nxcutlength, nxpowderlength, nxrapidlength, nxspnetarea, nxusedarea, nxused, nxmarkcount, nxprofcount, nxspischanged, nxspinsertdate
		FROM dbo.nxsheetpath
		WHERE nxpathid = @p1`
	logf(tenantID, sqlstr, id)
	n := SheetPathModel{}
	row := r.dbs[tenantID].QueryRowxContext(ctx, sqlstr, id)
	err := row.StructScan(&n)
	if err != nil {
		return nil, logerror(err)
	}
	return mapSheetPathModel(&n), nil
}

func (r *SheetPathRepository) GetByPathIDs(ctx context.Context, ids []int64) ([]*sheetpath.SheetPath, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT
		nxsheetpathid, nxpathid, nxmatorderlineid, nxinventoryid, nxnestreservationid, nxproductid, nxpartid, nxcutlength, nxpowderlength, nxrapidlength, nxspnetarea, nxusedarea, nxused, nxmarkcount, nxprofcount, nxspischanged, nxspinsertdate
		FROM dbo.nxsheetpath
		WHERE nxpathid IN (?)`
	query, args, err := sqlx.In(sqlstr, ids)
	if err != nil {
		return nil, logerror(err)
	}
	query = r.dbs[tenantID].Rebind(query)

	logf(tenantID, query, ids)

	rows, err := r.dbs[tenantID].QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, logerror(err)
	}

	var nn []*sheetpath.SheetPath
	for rows.Next() {
		var n SheetPathModel
		err := rows.StructScan(&n)
		if err != nil {
			return nil, logerror(err)
		}
		nn = append(nn, mapSheetPathModel(&n))
	}

	return nn, nil
}

func mapSheetPathModel(m *SheetPathModel) *sheetpath.SheetPath {
	var inventoryID *int64
	if m.InventoryID.Valid {
		inventoryID = &m.InventoryID.Int64
	}
	var productID *int64
	if m.ProductID.Valid {
		productID = &m.ProductID.Int64
	}

	sp := sheetpath.New(
		m.ID,
		m.PathID.Int64,
		m.MatOrderID,
		inventoryID,
		productID,
		m.NetArea.Float64,
		m.UsedArea.Float64,
		m.Used.Float64,
		m.InsertDate.Time,
	)
	return sp
}
