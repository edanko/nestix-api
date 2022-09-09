package adapters

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/edanko/nestix-api/internal/domain/machine"
	"github.com/edanko/nestix-api/pkg/tenant"
)

// MachineModel represents a row from 'dbo.machine'.
type MachineModel struct {
	ID          int64          `db:"machineid"`
	Name        string         `db:"name"`
	ControlType sql.NullString `db:"controltype"`
	Timestamp   sql.NullInt64  `db:"nxtimestamp"`
}

type MachineRepository struct {
	dbs map[string]*sqlx.DB
}

func NewMachineRepository(db map[string]*sqlx.DB) *MachineRepository {
	return &MachineRepository{
		dbs: db,
	}
}

// Insert inserts the MachineModel to the database.
func (r *MachineRepository) Insert(ctx context.Context, m *MachineModel) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `INSERT INTO dbo.machine (
		machineid, name, controltype, nxtimestamp
		) VALUES (
		@p1, @p2, @p3, @p4
		)`
	logf(tenantID, sqlstr, m.ID, m.Name, m.ControlType, m.Timestamp)
	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, m.ID, m.Name, m.ControlType, m.Timestamp); err != nil {
		return logerror(err)
	}

	return nil
}

// Update updates a MachineModel in the database.
func (r *MachineRepository) Update(ctx context.Context, m *MachineModel) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `UPDATE dbo.machine SET
		name = @p1, controltype = @p2, nxtimestamp = @p3
		WHERE machineid = @p4`
	logf(tenantID, sqlstr, m.Name, m.ControlType, m.Timestamp, m.ID)
	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, m.Name, m.ControlType, m.Timestamp, m.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Delete deletes the MachineModel from the database.
func (r *MachineRepository) Delete(ctx context.Context, id int64) error {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return errors.New("tenant id not found in context")
	}

	const sqlstr = `DELETE FROM dbo.machine
		WHERE machineid = @p1`
	logf(tenantID, sqlstr, id)
	if _, err := r.dbs[tenantID].ExecContext(ctx, sqlstr, id); err != nil {
		return logerror(err)
	}

	return nil
}

func (r *MachineRepository) List(ctx context.Context) ([]*machine.Machine, error) {
	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, errors.New("tenant id not found in context")
	}

	const sqlstr = `SELECT
		machineid, name, controltype, nxtimestamp
		FROM dbo.machine`
	logf(tenantID, sqlstr)

	rows, err := r.dbs[tenantID].QueryxContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	var nn []*machine.Machine
	for rows.Next() {
		var n MachineModel
		err := rows.StructScan(&n)
		if err != nil {
			return nil, logerror(err)
		}
		nn = append(nn, mapMachineEntity(&n))
	}

	return nn, nil
}

func mapMachineEntity(m *MachineModel) *machine.Machine {
	machine, err := machine.New(m.ID, m.Name)
	if err != nil {
		log.Error().Err(err).Msg("machine failed")
	}
	return machine
}
