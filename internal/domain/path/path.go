package path

import (
	"strings"
	"time"
)

type Path struct {
	id         int64
	name       string
	machineID  int64
	info       *string
	length     float64
	height     float64
	metafile   string
	site       string
	filename   string
	torchData  []byte
	bevelInfo  []byte
	toolData   []byte
	bevelData  []byte
	insertDate time.Time
	creator    string
	created    time.Time
	changer    *string
	changed    *time.Time
	changeType ChangeType
}

func (p *Path) ID() int64 {
	return p.id
}

func (p *Path) Name() string {
	return p.name
}

func (p *Path) MachineID() int64 {
	return p.machineID
}

func (p *Path) Info() *string {
	return p.info
}

func (p *Path) MainLength() float64 {
	return p.length
}

func (p *Path) MainHeight() float64 {
	return p.height
}

func (p *Path) Metafile() string {
	return p.metafile
}

func (p *Path) Site() string {
	return p.site
}

func (p *Path) Filename() string {
	return p.filename
}

func (p *Path) TorchData() []byte {
	return p.torchData
}

func (p *Path) BevelInfo() []byte {
	return p.bevelInfo
}

func (p *Path) ToolData() []byte {
	return p.toolData
}

func (p *Path) BevelData() []byte {
	return p.bevelData
}

func (p *Path) InsertDate() time.Time {
	return p.insertDate
}

func (p *Path) Creator() string {
	return p.creator
}

func (p *Path) Created() time.Time {
	return p.created
}

func (p *Path) Changer() *string {
	return p.changer
}

func (p *Path) Changed() *time.Time {
	return p.changed
}

func (p *Path) ChangeType() ChangeType {
	return p.changeType
}

func New(
	id int64,
	name string,
	machineID int64,
	info *string,
	length float64,
	height float64,
	metafile string,
	site string,
	filename string,
	torchData []byte,
	bevelInfo []byte,
	toolData []byte,
	bevelData []byte,
	insertDate time.Time,
	creator string,
	created time.Time,
	changer *string,
	changed *time.Time,
	changeType *int64,
) (*Path, error) {
	if length == 0 {
		return nil, ErrValidationLength
	}
	if height == 0 {
		return nil, ErrValidationHeight
	}
	if len(metafile) == 0 || !strings.HasSuffix(metafile, ".emf") {
		return nil, ErrValidationMetafile
	}
	if len(filename) == 0 || !strings.HasSuffix(filename, ".nest.nxl") {
		return nil, ErrValidationFilename
	}

	return &Path{
		id:         id,
		name:       name,
		machineID:  machineID,
		info:       info,
		length:     length,
		height:     height,
		metafile:   metafile,
		site:       site,
		filename:   filename,
		torchData:  torchData,
		bevelInfo:  bevelInfo,
		toolData:   toolData,
		bevelData:  bevelData,
		insertDate: insertDate,
		creator:    creator,
		created:    created,
		changer:    changer,
		changed:    changed,
		changeType: NewChangeTypeFromInt64(changeType),
	}, nil
}
