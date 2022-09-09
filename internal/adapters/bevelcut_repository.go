package adapters

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/edanko/nestix-api/pkg/tenant"
)

// BevelCutModel represents a row from 'dbo.nxbevelcut'.
type BevelCutModel struct {
	ID           int64           `db:"nxbevelcutid"`
	Name         string          `db:"nxbvlcutname"`
	AngleMin     sql.NullFloat64 `db:"nxbvlanglemin"`
	AngleMax     sql.NullFloat64 `db:"nxbvlanglemax"`
	RootWidthMin sql.NullFloat64 `db:"nxbvlrootwidthmin"`
	RootWidthMax sql.NullFloat64 `db:"nxbvlrootwidthmax"`
	EdgeDistMin  sql.NullFloat64 `db:"nxbvledgedistmin"`
	EdgeDistMax  sql.NullFloat64 `db:"nxbvledgedistmax"`
	ThickMin     float32         `db:"nxbvlthickmin"`
	ThickMax     float32         `db:"nxbvlthickmax"`
	Speed        float32         `db:"nxbvlspeed"`
	CorrAngle    float32         `db:"nxbvlcorrangle"`
	CorrWidth    float32         `db:"nxbvlcorrwidth"`
	Tech         sql.NullString  `db:"nxbvltech"`
	TechGroup    sql.NullString  `db:"nxbvltechgroup"`
	MatGroup     sql.NullString  `db:"nxbvlmatgroup"`
	Creator      sql.NullString  `db:"nxbvlcreator"`
	Created      sql.NullTime    `db:"nxbvlcreated"`
	Changer      sql.NullString  `db:"nxbvlchanger"`
	Changed      sql.NullTime    `db:"nxbvlchanged"`
}

type BevelCutRepository struct {
	dbs map[string]*sqlx.DB
}

func NewBevelCutRepository(db map[string]*sqlx.DB) *BevelCutRepository {
	return &BevelCutRepository{
		dbs: db,
	}
}

// Insert inserts the BevelCutModel to the database.
func (r *BevelCutRepository) Insert(ctx context.Context, n *BevelCutModel) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `INSERT INTO dbo.nxbevelcut (
		nxbevelcutid, nxbvlcutname, nxbvlanglemin, nxbvlanglemax, nxbvlrootwidthmin, nxbvlrootwidthmax, nxbvledgedistmin, nxbvledgedistmax, nxbvlthickmin, nxbvlthickmax, nxbvlspeed, nxbvlcorrangle, nxbvlcorrwidth, nxbvltech, nxbvltechgroup, nxbvlmatgroup, nxbvlcreator, nxbvlcreated, nxbvlchanger, nxbvlchanged
		) VALUES (
		@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, @p13, @p14, @p15, @p16, @p17, @p18, @p19, @p20
		)`
	logf(
		tenantID,
		sqlstr,
		n.ID,
		n.Name,
		n.AngleMin,
		n.AngleMax,
		n.RootWidthMin,
		n.RootWidthMax,
		n.EdgeDistMin,
		n.EdgeDistMax,
		n.ThickMin,
		n.ThickMax,
		n.Speed,
		n.CorrAngle,
		n.CorrWidth,
		n.Tech,
		n.TechGroup,
		n.MatGroup,
		n.Creator,
		n.Created,
		n.Changer,
		n.Changed,
	)
	if _, err := r.dbs[tenantID].ExecContext(
		ctx,
		sqlstr,
		n.ID,
		n.Name,
		n.AngleMin,
		n.AngleMax,
		n.RootWidthMin,
		n.RootWidthMax,
		n.EdgeDistMin,
		n.EdgeDistMax,
		n.ThickMin,
		n.ThickMax,
		n.Speed,
		n.CorrAngle,
		n.CorrWidth,
		n.Tech,
		n.TechGroup,
		n.MatGroup,
		n.Creator,
		n.Created,
		n.Changer,
		n.Changed,
	); err != nil {
		return logerror(err)
	}

	return nil
}

// Update updates a BevelCutModel in the database.
func (r *BevelCutRepository) Update(ctx context.Context, n *BevelCutModel) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `UPDATE dbo.nxbevelcut SET
		nxbvlcutname = @p1, nxbvlanglemin = @p2, nxbvlanglemax = @p3, nxbvlrootwidthmin = @p4, nxbvlrootwidthmax = @p5, nxbvledgedistmin = @p6, nxbvledgedistmax = @p7, nxbvlthickmin = @p8, nxbvlthickmax = @p9, nxbvlspeed = @p10, nxbvlcorrangle = @p11, nxbvlcorrwidth = @p12, nxbvltech = @p13, nxbvltechgroup = @p14, nxbvlmatgroup = @p15, nxbvlcreator = @p16, nxbvlcreated = @p17, nxbvlchanger = @p18, nxbvlchanged = @p19
		WHERE nxbevelcutid = @p20`
	logf(
		tenantID,
		sqlstr,
		n.Name,
		n.AngleMin,
		n.AngleMax,
		n.RootWidthMin,
		n.RootWidthMax,
		n.EdgeDistMin,
		n.EdgeDistMax,
		n.ThickMin,
		n.ThickMax,
		n.Speed,
		n.CorrAngle,
		n.CorrWidth,
		n.Tech,
		n.TechGroup,
		n.MatGroup,
		n.Creator,
		n.Created,
		n.Changer,
		n.Changed,
		n.ID,
	)
	if _, err := r.dbs[tenantID].ExecContext(
		ctx,
		sqlstr,
		n.Name,
		n.AngleMin,
		n.AngleMax,
		n.RootWidthMin,
		n.RootWidthMax,
		n.EdgeDistMin,
		n.EdgeDistMax,
		n.ThickMin,
		n.ThickMax,
		n.Speed,
		n.CorrAngle,
		n.CorrWidth,
		n.Tech,
		n.TechGroup,
		n.MatGroup,
		n.Creator,
		n.Created,
		n.Changer,
		n.Changed,
		n.ID,
	); err != nil {
		return logerror(err)
	}
	return nil
}

// Delete deletes the BevelCutModel from the database.
func (r *BevelCutRepository) Delete(ctx context.Context, id int64) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `DELETE FROM dbo.nxbevelcut ` +
		`WHERE nxbevelcutid = @p1`
	logf(tenantID, sqlstr, id)
	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, id); err != nil {
		return logerror(err)
	}

	return nil
}

func (r *BevelCutRepository) List(ctx context.Context) ([]*BevelCutModel, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT
		nxbevelcutid, nxbvlcutname, nxbvlanglemin, nxbvlanglemax, nxbvlrootwidthmin, nxbvlrootwidthmax, nxbvledgedistmin, nxbvledgedistmax, nxbvlthickmin, nxbvlthickmax, nxbvlspeed, nxbvlcorrangle, nxbvlcorrwidth, nxbvltech, nxbvltechgroup, nxbvlmatgroup, nxbvlcreator, nxbvlcreated, nxbvlchanger, nxbvlchanged
		FROM dbo.nxbevelcut`
	logf(tenantID, sqlstr)

	rows, err := r.dbs[tenantID].QueryxContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	var nn []*BevelCutModel
	for rows.Next() {
		var n BevelCutModel
		err := rows.StructScan(&n)
		if err != nil {
			return nil, logerror(err)
		}
		nn = append(nn, &n)
	}

	return nn, nil
}

// func mapEntity(e *BevelCutModel) *bevelcut.BevelCut {
// 	return bevelcut.NewBevelCut(
// 		e.ID,
// 		e.Name,
// 		e.AngleMin,
// 		e.AngleMax,
// 		e.RootWidthMin,
// 		e.RootWidthMax,
// 		e.EdgeDistMin,
// 		e.EdgeDistMax,
// 		e.ThickMin,
// 		e.ThickMax,
// 		e.Speed,
// 		e.CorrAngle,
// 		e.CorrWidth,
// 		e.Tech,
// 		e.TechGroup,
// 		e.materialGroup,
// 		e.creator,
// 		e.created,
// 		e.changer,
// 		e.changed,
// 	)
// }
