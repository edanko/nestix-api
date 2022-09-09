package bevelcut

import "time"

type BevelCut struct {
	id           int64
	name         string
	angleMin     *float64
	angleMax     *float64
	rootWidthMin *float64
	rootWidthMax *float64
	edgeDistMin  *float64
	edgeDistMax  *float64
	thickMin     float32
	thickMax     float32
	speed        float32
	corrAngle    float32
	corrWidth    float32
	tech         *string
	techGroup    *string
	matGroup     *string
	creator      *string
	created      *time.Time
	changer      *string
	changed      *time.Time
}

func NewBevelCut(
	id int64,
	name string,
	angleMin *float64,
	angleMax *float64,
	rootWidthMin *float64,
	rootWidthMax *float64,
	edgeDistMin *float64,
	edgeDistMax *float64,
	thickMin float32,
	thickMax float32,
	speed float32,
	corrAngle float32,
	corrWidth float32,
	tech *string,
	techGroup *string,
	matGroup *string,
	creator *string,
	created *time.Time,
	changer *string,
	changed *time.Time,
) *BevelCut {
	return &BevelCut{
		id:           id,
		name:         name,
		angleMin:     angleMin,
		angleMax:     angleMax,
		rootWidthMin: rootWidthMin,
		rootWidthMax: rootWidthMax,
		edgeDistMin:  edgeDistMin,
		edgeDistMax:  edgeDistMax,
		thickMin:     thickMin,
		thickMax:     thickMax,
		speed:        speed,
		corrAngle:    corrAngle,
		corrWidth:    corrWidth,
		tech:         tech,
		techGroup:    techGroup,
		matGroup:     matGroup,
		creator:      creator,
		created:      created,
		changer:      changer,
		changed:      changed,
	}
}
