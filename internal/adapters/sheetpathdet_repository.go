package adapters

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	sheetpathdet2 "github.com/edanko/nestix-api/internal/domain/sheetpathdet"
	"github.com/edanko/nestix-api/pkg/tenant"
)

// SheetPathDetModel represents a row from 'dbo.nxsheetpathdet'.
type SheetPathDetModel struct {
	ID          int64           `db:"nxsheetpathdetid"`
	SheetPathID int64           `db:"nxsheetpathid"`
	OrderID     sql.NullInt64   `db:"nxorderlineid"`
	PartID      sql.NullInt64   `db:"nxpartid"`
	ProductID   sql.NullInt64   `db:"nxproductid"`
	DetailCount sql.NullInt64   `db:"nxdetailcount"`
	DetailCode  sql.NullString  `db:"nxdetailcode"`
	UsedArea    sql.NullFloat64 `db:"nxusedarea"`
	Area        sql.NullFloat64 `db:"nxarea"`
	SlagArea    sql.NullFloat64 `db:"nxslagarea"`
	NestScrap   sql.NullFloat64 `db:"nxspdnestscrap"`
	AddScrap    sql.NullFloat64 `db:"nxspdaddscrap"`
	MachTime    sql.NullFloat64 `db:"nxspdmachtime"`
	ClaimArea   sql.NullFloat64 `db:"nxspdclaimarea"`
	ReturnArea  sql.NullFloat64 `db:"nxspdreturnarea"`
	MachTimePct sql.NullFloat64 `db:"nxspdmachtimepct"`
	DetailType  sql.NullInt64   `db:"nxspdtype"`
	SequenceNo  sql.NullInt64   `db:"nxspdsequenceno"`
}

type SheetPathDetRepository struct {
	dbs map[string]*sqlx.DB
}

func NewSheetPathDetRepository(db map[string]*sqlx.DB) *SheetPathDetRepository {
	return &SheetPathDetRepository{
		dbs: db,
	}
}

// Insert inserts the SheetPathDetModel to the database.
func (r *SheetPathDetRepository) Insert(ctx context.Context, n *SheetPathDetModel) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `INSERT INTO dbo.nxsheetpathdet (
		nxsheetpathdetid, nxsheetpathid, nxorderlineid, nxpartid, nxproductid, nxdetailcount, nxdetailcode, nxusedarea, nxarea, nxslagarea, nxspdnestscrap, nxspdaddscrap, nxspdmachtime, nxspdclaimarea, nxspdreturnarea, nxspdmachtimepct, nxspdtype, nxspdsequenceno
		) VALUES (
		@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, @p13, @p14, @p15, @p16, @p17, @p18
		)`
	logf(
		tenantID,
		sqlstr,
		n.ID,
		n.SheetPathID,
		n.OrderID,
		n.PartID,
		n.ProductID,
		n.DetailCount,
		n.DetailCode,
		n.UsedArea,
		n.Area,
		n.SlagArea,
		n.NestScrap,
		n.AddScrap,
		n.MachTime,
		n.ClaimArea,
		n.ReturnArea,
		n.MachTimePct,
		n.DetailType,
		n.SequenceNo,
	)
	if _, err := r.dbs[tenantID].ExecContext(
		ctx,
		sqlstr,
		n.ID,
		n.SheetPathID,
		n.OrderID,
		n.PartID,
		n.ProductID,
		n.DetailCount,
		n.DetailCode,
		n.UsedArea,
		n.Area,
		n.SlagArea,
		n.NestScrap,
		n.AddScrap,
		n.MachTime,
		n.ClaimArea,
		n.ReturnArea,
		n.MachTimePct,
		n.DetailType,
		n.SequenceNo,
	); err != nil {
		return logerror(err)
	}
	return nil
}

// Update updates a SheetPathDetModel in the database.
func (r *SheetPathDetRepository) Update(ctx context.Context, n *SheetPathDetModel) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `UPDATE dbo.nxsheetpathdet SET
		nxsheetpathid = @p1, nxorderlineid = @p2, nxpartid = @p3, nxproductid = @p4, nxdetailcount = @p5, nxdetailcode = @p6, nxusedarea = @p7, nxarea = @p8, nxslagarea = @p9, nxspdnestscrap = @p10, nxspdaddscrap = @p11, nxspdmachtime = @p12, nxspdclaimarea = @p13, nxspdreturnarea = @p14, nxspdmachtimepct = @p15, nxspdtype = @p16, nxspdsequenceno = @p17
		WHERE nxsheetpathdetid = @p18`
	logf(
		tenantID,
		sqlstr,
		n.SheetPathID,
		n.OrderID,
		n.PartID,
		n.ProductID,
		n.DetailCount,
		n.DetailCode,
		n.UsedArea,
		n.Area,
		n.SlagArea,
		n.NestScrap,
		n.AddScrap,
		n.MachTime,
		n.ClaimArea,
		n.ReturnArea,
		n.MachTimePct,
		n.DetailType,
		n.SequenceNo,
		n.ID,
	)
	if _, err := r.dbs[tenantID].ExecContext(
		ctx,
		sqlstr,
		n.SheetPathID,
		n.OrderID,
		n.PartID,
		n.ProductID,
		n.DetailCount,
		n.DetailCode,
		n.UsedArea,
		n.Area,
		n.SlagArea,
		n.NestScrap,
		n.AddScrap,
		n.MachTime,
		n.ClaimArea,
		n.ReturnArea,
		n.MachTimePct,
		n.DetailType,
		n.SequenceNo,
		n.ID,
	); err != nil {
		return logerror(err)
	}
	return nil
}

// Delete deletes the SheetPathDetModel from the database.
func (r *SheetPathDetRepository) Delete(ctx context.Context, n *SheetPathDetModel) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `DELETE FROM dbo.nxsheetpathdet
		WHERE nxsheetpathdetid = @p1`
	logf(tenantID, sqlstr, n.ID)
	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, n.ID); err != nil {
		return logerror(err)
	}
	return nil
}

func (r *SheetPathDetRepository) ListPartsBySheetPathID(ctx context.Context, id int64) ([]*sheetpathdet2.Part, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT
		nxsheetpathdetid, nxsheetpathid, nxorderlineid, nxpartid, nxproductid, nxdetailcount, nxdetailcode, nxusedarea, nxarea, nxslagarea, nxspdnestscrap, nxspdaddscrap, nxspdmachtime, nxspdclaimarea, nxspdreturnarea, nxspdmachtimepct, nxspdtype, nxspdsequenceno
		FROM dbo.nxsheetpathdet
		WHERE nxspdtype = 1 AND nxsheetpathid = @p1
		ORDER BY nxdetailcode ASC`
	logf(tenantID, sqlstr, id)

	rows, err := r.dbs[tenantID].QueryxContext(ctx, sqlstr, id)
	if err != nil {
		return nil, logerror(err)
	}

	var nn []*sheetpathdet2.Part

	for rows.Next() {
		n := SheetPathDetModel{}

		err := rows.StructScan(&n)
		if err != nil {
			return nil, logerror(err)
		}
		nn = append(nn, mapSheetPathDetPartModel(&n))
	}
	return nn, nil
}

func (r *SheetPathDetRepository) ListRemnantsBySheetPathID(ctx context.Context, id int64) ([]*sheetpathdet2.Remnant, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT
		nxsheetpathdetid, nxsheetpathid, nxorderlineid, nxpartid, nxproductid, nxdetailcount, nxdetailcode, nxusedarea, nxarea, nxslagarea, nxspdnestscrap, nxspdaddscrap, nxspdmachtime, nxspdclaimarea, nxspdreturnarea, nxspdmachtimepct, nxspdtype, nxspdsequenceno
		FROM dbo.nxsheetpathdet
		WHERE nxspdtype = 2 AND nxsheetpathid = @p1`
	logf(tenantID, sqlstr, id)

	rows, err := r.dbs[tenantID].QueryxContext(ctx, sqlstr, id)
	if err != nil {
		return nil, logerror(err)
	}

	var nn []*sheetpathdet2.Remnant

	for rows.Next() {
		n := SheetPathDetModel{}

		err := rows.StructScan(&n)
		if err != nil {
			return nil, logerror(err)
		}
		nn = append(nn, mapSheetPathDetRemnantModel(&n))
	}
	return nn, nil
}

func mapSheetPathDetPartModel(m *SheetPathDetModel) *sheetpathdet2.Part {
	var orderID *int64
	if m.OrderID.Valid {
		orderID = &m.OrderID.Int64
	}
	var partID *int64
	if m.PartID.Valid {
		partID = &m.PartID.Int64
	}

	d, err := sheetpathdet2.NewPart(
		m.ID,
		m.SheetPathID,
		orderID,
		partID,
		m.DetailCount.Int64,
		m.DetailCode.String,
		m.UsedArea.Float64,
		m.Area.Float64,
		m.SlagArea.Float64,
		m.NestScrap.Float64,
		m.ClaimArea.Float64,
		m.ReturnArea.Float64,
		m.SequenceNo.Int64,
	)

	if err != nil {
		log.Error().Err(err).Msg("mapSheetPathDetModel")
	}
	return d
}

func mapSheetPathDetRemnantModel(m *SheetPathDetModel) *sheetpathdet2.Remnant {
	var productID *int64
	if m.ProductID.Valid {
		productID = &m.ProductID.Int64
	}

	d, err := sheetpathdet2.NewRemnant(
		m.ID,
		m.SheetPathID,
		productID,
		m.DetailCount.Int64,
		m.Area.Float64,
		m.SequenceNo.Int64,
	)

	if err != nil {
		log.Error().Err(err).Msg("mapSheetPathDetRemnantModel")
	}
	return d
}
