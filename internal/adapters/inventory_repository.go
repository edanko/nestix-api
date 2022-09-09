package adapters

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/edanko/nestix-api/pkg/tenant"
)

// InventoryModel represents a row from 'dbo.nxinventory'.
type InventoryModel struct {
	ID             int64           `db:"nxinventoryid"`
	PartID         sql.NullInt64   `db:"nxpartid"`
	ProductID      sql.NullInt64   `db:"nxproductid"`
	MatOrderID     sql.NullInt64   `db:"nxinvmatorderlineid"`
	Name           sql.NullString  `db:"nxinvname"`
	Charge         sql.NullString  `db:"nxinvcharge"`
	Length         sql.NullFloat64 `db:"nxinvlength"`
	Width          sql.NullFloat64 `db:"nxinvwidth"`
	Thick          sql.NullFloat64 `db:"nxinvthick"`
	Weight         sql.NullFloat64 `db:"nxinvweight"`
	Area           sql.NullFloat64 `db:"nxinvarea"`
	PlateNo        sql.NullString  `db:"nxinvplateno"`
	Type           sql.NullInt64   `db:"nxinvtype"`
	Count          sql.NullInt64   `db:"nxinvcount"`
	Info           sql.NullString  `db:"nxinvinfotxt"`
	BatchNo        sql.NullString  `db:"nxinvbatchno"`
	ActDate        sql.NullTime    `db:"nxinvactdate"`
	ActCode        sql.NullInt64   `db:"nxinvactcode"`
	ActTxt         sql.NullString  `db:"nxinvacttxt"`
	MasterID       sql.NullInt64   `db:"nxinvmasterid"`
	Quality        sql.NullString  `db:"nxinvquality"`
	Classification sql.NullString  `db:"nxinvclassification"`
	CertNo         sql.NullString  `db:"nxinvcertno"`
	CertType       sql.NullString  `db:"nxinvcerttype"`
	RawMatID       sql.NullInt64   `db:"nxinvrawmatid"`
	OrderID        sql.NullInt64   `db:"nxorderlineid"`
	InsertDate     sql.NullTime    `db:"nxinvinsertdate"`
	SheetPathDetID sql.NullInt64   `db:"nxsheetpathdetid"`
	Source         sql.NullInt64   `db:"nxinvsource"`
	OrderNo        sql.NullString  `db:"nxinvorderno"`
}

type InventoryRepository struct {
	dbs map[string]*sqlx.DB
}

func NewInventoryRepository(db map[string]*sqlx.DB) *InventoryRepository {
	return &InventoryRepository{
		dbs: db,
	}
}

// Insert inserts the InventoryModel to the database.
func (r *InventoryRepository) Insert(ctx context.Context, n *InventoryModel) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `INSERT INTO dbo.nxinventory (
		nxinventoryid, nxpartid, nxproductid, nxinvmatorderlineid, nxinvname, nxinvcharge, nxinvlength, nxinvwidth, nxinvthick, nxinvweight, nxinvarea, nxinvplateno, nxinvtype, nxinvcount, nxinvinfotxt, nxinvbatchno, nxinvactdate, nxinvactcode, nxinvacttxt, nxinvmasterid, nxinvquality, nxinvclassification, nxinvcertno, nxinvcerttype, nxinvrawmatid, nxorderlineid, nxinvinsertdate, nxsheetpathdetid, nxinvsource, nxinvorderno
		) VALUES (
		@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, @p13, @p14, @p15, @p16, @p17, @p18, @p19, @p20, @p21, @p22, @p23, @p24, @p25, @p26, @p27, @p28, @p29, @p30
		)`
	logf(
		tenantID,
		sqlstr,
		n.ID,
		n.PartID,
		n.ProductID,
		n.MatOrderID,
		n.Name,
		n.Charge,
		n.Length,
		n.Width,
		n.Thick,
		n.Weight,
		n.Area,
		n.PlateNo,
		n.Type,
		n.Count,
		n.Info,
		n.BatchNo,
		n.ActDate,
		n.ActCode,
		n.ActTxt,
		n.MasterID,
		n.Quality,
		n.Classification,
		n.CertNo,
		n.CertType,
		n.RawMatID,
		n.OrderID,
		n.InsertDate,
		n.SheetPathDetID,
		n.Source,
		n.OrderNo,
	)
	if _, err := r.dbs[tenantID].ExecContext(
		ctx,
		sqlstr,
		n.ID,
		n.PartID,
		n.ProductID,
		n.MatOrderID,
		n.Name,
		n.Charge,
		n.Length,
		n.Width,
		n.Thick,
		n.Weight,
		n.Area,
		n.PlateNo,
		n.Type,
		n.Count,
		n.Info,
		n.BatchNo,
		n.ActDate,
		n.ActCode,
		n.ActTxt,
		n.MasterID,
		n.Quality,
		n.Classification,
		n.CertNo,
		n.CertType,
		n.RawMatID,
		n.OrderID,
		n.InsertDate,
		n.SheetPathDetID,
		n.Source,
		n.OrderNo,
	); err != nil {
		return logerror(err)
	}

	return nil
}

// Update updates an InventoryModel in the database.
func (r *InventoryRepository) Update(ctx context.Context, n *InventoryModel) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `UPDATE dbo.nxinventory SET
		nxpartid = @p1, nxproductid = @p2, nxinvmatorderlineid = @p3, nxinvname = @p4, nxinvcharge = @p5, nxinvlength = @p6, nxinvwidth = @p7, nxinvthick = @p8, nxinvweight = @p9, nxinvarea = @p10, nxinvplateno = @p11, nxinvtype = @p12, nxinvcount = @p13, nxinvinfotxt = @p14, nxinvbatchno = @p15, nxinvactdate = @p16, nxinvactcode = @p17, nxinvacttxt = @p18, nxinvmasterid = @p19, nxinvquality = @p20, nxinvclassification = @p21, nxinvcertno = @p22, nxinvcerttype = @p23, nxinvrawmatid = @p24, nxorderlineid = @p25, nxinvinsertdate = @p26, nxsheetpathdetid = @p27, nxinvsource = @p28, nxinvorderno = @p29
		WHERE nxinventoryid = @p30`
	logf(
		tenantID,
		sqlstr,
		n.PartID,
		n.ProductID,
		n.MatOrderID,
		n.Name,
		n.Charge,
		n.Length,
		n.Width,
		n.Thick,
		n.Weight,
		n.Area,
		n.PlateNo,
		n.Type,
		n.Count,
		n.Info,
		n.BatchNo,
		n.ActDate,
		n.ActCode,
		n.ActTxt,
		n.MasterID,
		n.Quality,
		n.Classification,
		n.CertNo,
		n.CertType,
		n.RawMatID,
		n.OrderID,
		n.InsertDate,
		n.SheetPathDetID,
		n.Source,
		n.OrderNo,
		n.ID,
	)
	if _, err := r.dbs[tenantID].ExecContext(
		ctx,
		sqlstr,
		n.PartID,
		n.ProductID,
		n.MatOrderID,
		n.Name,
		n.Charge,
		n.Length,
		n.Width,
		n.Thick,
		n.Weight,
		n.Area,
		n.PlateNo,
		n.Type,
		n.Count,
		n.Info,
		n.BatchNo,
		n.ActDate,
		n.ActCode,
		n.ActTxt,
		n.MasterID,
		n.Quality,
		n.Classification,
		n.CertNo,
		n.CertType,
		n.RawMatID,
		n.OrderID,
		n.InsertDate,
		n.SheetPathDetID,
		n.Source,
		n.OrderNo,
		n.ID,
	); err != nil {
		return logerror(err)
	}
	return nil
}

// Delete deletes the InventoryModel from the database.
func (r *InventoryRepository) Delete(ctx context.Context, id int64) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `DELETE FROM dbo.nxinventory ` +
		`WHERE nxinventoryid = @p1`
	logf(tenantID, sqlstr, id)
	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, id); err != nil {
		return logerror(err)
	}

	return nil
}
