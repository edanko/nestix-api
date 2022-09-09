package adapters

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"

	"github.com/edanko/nestix-api/internal/domain/path"
	"github.com/edanko/nestix-api/pkg/tenant"
)

// PathModel represents a row from 'dbo.nxpath'.
type PathModel struct {
	ID              int64           `db:"nxpathid"`
	Name            string          `db:"nxname"`
	MachineID       int64           `db:"nxmachineid"`
	CutLength       sql.NullFloat64 `db:"nxcutlength"`
	PowderLength    sql.NullFloat64 `db:"nxpowderlength"`
	RapidLength     sql.NullFloat64 `db:"nxrapidlength"`
	MarkCount       sql.NullInt64   `db:"nxmarkcount"`
	ProfCount       sql.NullInt64   `db:"nxprofcount"`
	Info            sql.NullString  `db:"nxpathinfo"`
	Length          sql.NullFloat64 `db:"nxmainlength"`
	Height          sql.NullFloat64 `db:"nxmainheight"`
	Metafile        sql.NullString  `db:"nxpthmetafile"`
	MachTime        sql.NullFloat64 `db:"nxmachtime"`
	Site            sql.NullString  `db:"nxpthsite"`
	Filename        sql.NullString  `db:"nxpthfilename"`
	TorchData       []byte          `db:"nxpthtorchdata"`
	BevelInfo       []byte          `db:"nxpthbevelinfo"`
	ToolData        []byte          `db:"nxpthtooldata"`
	BevelData       []byte          `db:"nxpthbeveldata"`
	InsertDate      time.Time       `db:"nxpthinsertdate"`
	Creator         sql.NullString  `db:"nxpthcreator"`
	Created         sql.NullTime    `db:"nxpthcreated"`
	Changer         sql.NullString  `db:"nxpthchanger"`
	Changed         sql.NullTime    `db:"nxpthchanged"`
	Reusable        sql.NullInt64   `db:"nxpthreusable"`
	ChangeType      sql.NullInt64   `db:"nxpthchangetype"`
	SectionCopyData []byte          `db:"nxpthsectioncopydata"`
	Info1           sql.NullString  `db:"nxpthpathinfo1"`
	Info2           sql.NullString  `db:"nxpthpathinfo2"`
	Info3           sql.NullString  `db:"nxpthpathinfo3"`
	NestingData     []byte          `db:"nxpthnestingdata"`
}

type PathRepository struct {
	dbs map[string]*sqlx.DB
}

func NewPathRepository(db map[string]*sqlx.DB) *PathRepository {
	return &PathRepository{
		dbs: db,
	}
}

// Insert inserts the PathModel to the database.
func (r *PathRepository) Insert(ctx context.Context, n *PathModel) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `INSERT INTO dbo.nxpath (
		nxpathid, nxname, nxmachineid, nxcutlength, nxpowderlength, nxrapidlength, nxmarkcount, nxprofcount, nxpathinfo, nxmainlength, nxmainheight, nxpthmetafile, nxmachtime, nxpthsite, nxpthfilename, nxpthtorchdata, nxpthbevelinfo, nxpthtooldata, nxpthbeveldata, nxpthinsertdate, nxpthcreator, nxpthcreated, nxpthchanger, nxpthchanged, nxpthreusable, nxpthchangetype, nxpthsectioncopydata, nxpthpathinfo1, nxpthpathinfo2, nxpthpathinfo3, nxpthnestingdata
		) VALUES (
		@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, @p13, @p14, @p15, @p16, @p17, @p18, @p19, @p20, @p21, @p22, @p23, @p24, @p25, @p26, @p27, @p28, @p29, @p30, @p31
		)`
	logf(tenantID,
		sqlstr,
		n.ID,
		n.Name,
		n.MachineID,
		n.CutLength,
		n.PowderLength,
		n.RapidLength,
		n.MarkCount,
		n.ProfCount,
		n.Info,
		n.Length,
		n.Height,
		n.Metafile,
		n.MachTime,
		n.Site,
		n.Filename,
		n.TorchData,
		n.BevelInfo,
		n.ToolData,
		n.BevelData,
		n.InsertDate,
		n.Creator,
		n.Created,
		n.Changer,
		n.Changed,
		n.Reusable,
		n.ChangeType,
		n.SectionCopyData,
		n.Info1,
		n.Info2,
		n.Info3,
		n.NestingData,
	)
	if _, err := r.dbs[tenantID].ExecContext(
		ctx,
		sqlstr,
		n.ID,
		n.Name,
		n.MachineID,
		n.CutLength,
		n.PowderLength,
		n.RapidLength,
		n.MarkCount,
		n.ProfCount,
		n.Info,
		n.Length,
		n.Height,
		n.Metafile,
		n.MachTime,
		n.Site,
		n.Filename,
		n.TorchData,
		n.BevelInfo,
		n.ToolData,
		n.BevelData,
		n.InsertDate,
		n.Creator,
		n.Created,
		n.Changer,
		n.Changed,
		n.Reusable,
		n.ChangeType,
		n.SectionCopyData,
		n.Info1,
		n.Info2,
		n.Info3,
		n.NestingData,
	); err != nil {
		return logerror(err)
	}
	return nil
}

// Update updates a PathModel in the database.
func (r *PathRepository) Update(ctx context.Context, n *PathModel) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `UPDATE dbo.nxpath SET
		nxname = @p1, nxmachineid = @p2, nxcutlength = @p3, nxpowderlength = @p4, nxrapidlength = @p5, nxmarkcount = @p6, nxprofcount = @p7, nxpathinfo = @p8, nxmainlength = @p9, nxmainheight = @p10, nxpthmetafile = @p11, nxmachtime = @p12, nxpthsite = @p13, nxpthfilename = @p14, nxpthtorchdata = @p15, nxpthbevelinfo = @p16, nxpthtooldata = @p17, nxpthbeveldata = @p18, nxpthinsertdate = @p19, nxpthcreator = @p20, nxpthcreated = @p21, nxpthchanger = @p22, nxpthchanged = @p23, nxpthreusable = @p24, nxpthchangetype = @p25, nxpthsectioncopydata = @p26, nxpthpathinfo1 = @p27, nxpthpathinfo2 = @p28, nxpthpathinfo3 = @p29, nxpthnestingdata = @p30
		WHERE nxpathid = @p31`
	logf(
		tenantID,
		sqlstr,
		n.Name,
		n.MachineID,
		n.CutLength,
		n.PowderLength,
		n.RapidLength,
		n.MarkCount,
		n.ProfCount,
		n.Info,
		n.Length,
		n.Height,
		n.Metafile,
		n.MachTime,
		n.Site,
		n.Filename,
		n.TorchData,
		n.BevelInfo,
		n.ToolData,
		n.BevelData,
		n.InsertDate,
		n.Creator,
		n.Created,
		n.Changer,
		n.Changed,
		n.Reusable,
		n.ChangeType,
		n.SectionCopyData,
		n.Info1,
		n.Info2,
		n.Info3,
		n.NestingData,
		n.ID,
	)
	if _, err := r.dbs[tenantID].ExecContext(
		ctx,
		sqlstr,
		n.Name,
		n.MachineID,
		n.CutLength,
		n.PowderLength,
		n.RapidLength,
		n.MarkCount,
		n.ProfCount,
		n.Info,
		n.Length,
		n.Height,
		n.Metafile,
		n.MachTime,
		n.Site,
		n.Filename,
		n.TorchData,
		n.BevelInfo,
		n.ToolData,
		n.BevelData,
		n.InsertDate,
		n.Creator,
		n.Created,
		n.Changer,
		n.Changed,
		n.Reusable,
		n.ChangeType,
		n.SectionCopyData,
		n.Info1,
		n.Info2,
		n.Info3,
		n.NestingData,
		n.ID,
	); err != nil {
		return logerror(err)
	}
	return nil
}

// Delete deletes the PathModel from the database.
func (r *PathRepository) Delete(ctx context.Context, id int64) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const query = `DELETE FROM dbo.nxpath
		WHERE nxpathid = @p1`
	logf(tenantID, query, id)
	if _, err := r.dbs[tenantID].ExecContext(ctx, query, id); err != nil {
		return logerror(err)
	}
	return nil
}

// GetByID retrieves a row from 'dbo.nxpath' as a PathModel.
func (r *PathRepository) GetByID(ctx context.Context, id int64) (*path.Path, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const query = `SELECT
		nxpathid, nxname, nxmachineid, nxcutlength, nxpowderlength, nxrapidlength, nxmarkcount, nxprofcount, nxpathinfo, nxmainlength, nxmainheight, nxpthmetafile, nxmachtime, nxpthsite, nxpthfilename, nxpthtorchdata, nxpthbevelinfo, nxpthtooldata, nxpthbeveldata, nxpthinsertdate, nxpthcreator, nxpthcreated, nxpthchanger, nxpthchanged, nxpthreusable, nxpthchangetype, nxpthsectioncopydata, nxpthpathinfo1, nxpthpathinfo2, nxpthpathinfo3, nxpthnestingdata
		FROM dbo.nxpath
		WHERE nxpathid = @p1`
	logf(tenantID, query, id)

	var n PathModel
	row := r.dbs[tenantID].QueryRowxContext(ctx, query, id)
	err := row.StructScan(&n)
	if err != nil {
		return nil, logerror(err)
	}
	return mapPathModel(&n), nil
}

func (r *PathRepository) GetByIDs(ctx context.Context, ids []int64) ([]*path.Path, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	batches := lo.Chunk(ids, 32767)

	const sqlstr = `SELECT
		nxpathid, nxname, nxmachineid, nxcutlength, nxpowderlength, nxrapidlength, nxmarkcount, nxprofcount, nxpathinfo, nxmainlength, nxmainheight, nxpthmetafile, nxmachtime, nxpthsite, nxpthfilename, nxpthtorchdata, nxpthbevelinfo, nxpthtooldata, nxpthbeveldata, nxpthinsertdate, nxpthcreator, nxpthcreated, nxpthchanger, nxpthchanged, nxpthreusable, nxpthchangetype, nxpthsectioncopydata, nxpthpathinfo1, nxpthpathinfo2, nxpthpathinfo3, nxpthnestingdata
		FROM dbo.nxpath
		WHERE nxpathid IN (?)`

	var nn []*path.Path
	for _, batch := range batches {
		query, args, err := sqlx.In(sqlstr, batch)
		if err != nil {
			return nil, logerror(err)
		}
		query = r.dbs[tenantID].Rebind(query)

		logf(tenantID, query, args)

		rows, err := r.dbs[tenantID].QueryxContext(ctx, query, args...)
		if err != nil {
			return nil, logerror(err)
		}

		for rows.Next() {
			var n PathModel
			err := rows.StructScan(&n)
			if err != nil {
				return nil, logerror(err)
			}
			nn = append(nn, mapPathModel(&n))
		}
	}
	return nn, nil
}

// GetByName retrieves a row from 'dbo.nxpath' as a PathModel.
func (r *PathRepository) GetByName(ctx context.Context, name string) (*path.Path, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const query = `SELECT
		nxpathid, nxname, nxmachineid, nxcutlength, nxpowderlength, nxrapidlength, nxmarkcount, nxprofcount, nxpathinfo, nxmainlength, nxmainheight, nxpthmetafile, nxmachtime, nxpthsite, nxpthfilename, nxpthtorchdata, nxpthbevelinfo, nxpthtooldata, nxpthbeveldata, nxpthinsertdate, nxpthcreator, nxpthcreated, nxpthchanger, nxpthchanged, nxpthreusable, nxpthchangetype, nxpthsectioncopydata, nxpthpathinfo1, nxpthpathinfo2, nxpthpathinfo3, nxpthnestingdata
		FROM dbo.nxpath
		WHERE nxname = @p1`
	logf(tenantID, query, name)
	var n PathModel
	row := r.dbs[tenantID].QueryRowxContext(ctx, query, name)
	err := row.StructScan(&n)
	if err != nil {
		return nil, logerror(err)
	}
	return mapPathModel(&n), nil
}

func (r *PathRepository) List(ctx context.Context) ([]*path.Path, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const query = `SELECT
		nxpathid, nxname, nxmachineid, nxcutlength, nxpowderlength, nxrapidlength, nxmarkcount, nxprofcount, nxpathinfo, nxmainlength, nxmainheight, nxpthmetafile, nxmachtime, nxpthsite, nxpthfilename, nxpthtorchdata, nxpthbevelinfo, nxpthtooldata, nxpthbeveldata, nxpthinsertdate, nxpthcreator, nxpthcreated, nxpthchanger, nxpthchanged, nxpthreusable, nxpthchangetype, nxpthsectioncopydata, nxpthpathinfo1, nxpthpathinfo2, nxpthpathinfo3, nxpthnestingdata
		FROM dbo.nxpath`
	logf(tenantID, query)

	rows, err := r.dbs[tenantID].QueryxContext(ctx, query)
	if err != nil {
		return nil, logerror(err)
	}

	var nn []*path.Path
	for rows.Next() {
		var n PathModel
		err := rows.StructScan(&n)
		if err != nil {
			return nil, logerror(err)
		}
		nn = append(nn, mapPathModel(&n))
	}

	return nn, nil
}

func (r *PathRepository) ListInvalid(ctx context.Context) ([]*PathModel, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const query = `SELECT
		nxpathid, nxname, nxmachineid, nxcutlength, nxpowderlength, nxrapidlength, nxmarkcount, nxprofcount, nxpathinfo, nxmainlength, nxmainheight, nxpthmetafile, nxmachtime, nxpthsite, nxpthfilename, nxpthtorchdata, nxpthbevelinfo, nxpthtooldata, nxpthbeveldata, nxpthinsertdate, nxpthcreator, nxpthcreated, nxpthchanger, nxpthchanged, nxpthreusable, nxpthchangetype, nxpthsectioncopydata, nxpthpathinfo1, nxpthpathinfo2, nxpthpathinfo3, nxpthnestingdata
		FROM dbo.nxpath
		WHERE nxpthfilename IS NULL`
	logf(tenantID, query)

	rows, err := r.dbs[tenantID].QueryxContext(ctx, query)
	if err != nil {
		return nil, logerror(err)
	}

	var nn []*PathModel
	for rows.Next() {
		var n PathModel
		err := rows.StructScan(&n)
		if err != nil {
			return nil, logerror(err)
		}
		nn = append(nn, &n)
	}

	return nn, nil
}

func (r *PathRepository) SearchByName(ctx context.Context, name string) ([]*path.Path, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const query = `SELECT
		nxpathid, nxname, nxmachineid, nxcutlength, nxpowderlength, nxrapidlength, nxmarkcount, nxprofcount, nxpathinfo, nxmainlength, nxmainheight, nxpthmetafile, nxmachtime, nxpthsite, nxpthfilename, nxpthtorchdata, nxpthbevelinfo, nxpthtooldata, nxpthbeveldata, nxpthinsertdate, nxpthcreator, nxpthcreated, nxpthchanger, nxpthchanged, nxpthreusable, nxpthchangetype, nxpthsectioncopydata, nxpthpathinfo1, nxpthpathinfo2, nxpthpathinfo3, nxpthnestingdata
		FROM dbo.nxpath
		WHERE nxname LIKE '%'+@p1+'%'`
	logf(tenantID, query, name)

	rows, err := r.dbs[tenantID].QueryxContext(ctx, query, name)
	if err != nil {
		return nil, logerror(err)
	}

	var nn []*path.Path
	for rows.Next() {
		var n PathModel
		err := rows.StructScan(&n)
		if err != nil {
			return nil, logerror(err)
		}
		nn = append(nn, mapPathModel(&n))
	}

	return nn, nil
}

func mapPathModel(m *PathModel) *path.Path {
	var info *string
	if m.Info.Valid {
		info = &m.Info.String
	}
	var changer *string
	if m.Changer.Valid {
		changer = &m.Changer.String
	}
	var changed *time.Time
	if m.Changed.Valid {
		changed = &m.Changed.Time
	}
	var changeType *int64
	if m.ChangeType.Valid {
		changeType = &m.ChangeType.Int64
	}

	p, err := path.New(
		m.ID,
		m.Name,
		m.MachineID,
		info,
		m.Length.Float64,
		m.Height.Float64,
		m.Metafile.String,
		m.Site.String,
		m.Filename.String,
		m.TorchData,
		m.BevelInfo,
		m.TorchData,
		m.BevelData,
		m.InsertDate,
		m.Creator.String,
		m.Created.Time,
		changer,
		changed,
		changeType,
	)
	if err != nil {
		log.Error().Err(err).Msg("map path error")
	}

	return p
}
