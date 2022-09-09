package product

import "time"

type Plate struct {
	id            int64
	productNo     string
	length        float64
	width         float64
	thickness     float64
	materialGroup string
	quality       string
	density       float64
	insertDate    time.Time
	creator       string
	created       time.Time
	// drawn      int64 // always zero
}

func (p *Plate) ID() int64 {
	return p.id
}

func (p *Plate) ProductNo() string {
	return p.productNo
}

func (p *Plate) Length() float64 {
	return p.length
}

func (p *Plate) Width() float64 {
	return p.width
}

func (p *Plate) Thickness() float64 {
	return p.thickness
}

func (p *Plate) MaterialGroup() string {
	return p.materialGroup
}

func (p *Plate) Quality() string {
	return p.quality
}

func (p *Plate) Density() float64 {
	return p.density
}

func (p *Plate) InsertDate() time.Time {
	return p.insertDate
}

func (p *Plate) Creator() string {
	return p.creator
}

func (p *Plate) Created() time.Time {
	return p.created
}

func NewPlate(
	id int64,
	productNo string,
	length float64,
	width float64,
	thickness float64,
	materialGroup string,
	quality string,
	density float64,
	insertDate time.Time,
	creator string,
	created time.Time,
) *Plate {
	return &Plate{
		id:            id,
		productNo:     productNo,
		length:        length,
		width:         width,
		thickness:     thickness,
		materialGroup: materialGroup,
		quality:       quality,
		density:       density,
		insertDate:    insertDate,
		creator:       creator,
		created:       created,
	}
}
