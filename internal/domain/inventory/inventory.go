package inventory

import "time"

type Inventory struct {
	ID              int64
	PartID          *int64
	ProductID       *int64
	MaterialOrderID *int64
	Length          float64
	Width           float64
	Thick           *float64
	Weight          *float64
	Area            float64
	PlateNo         *string
	Type            *int64
	Count           int64
	ActCode         int64 // 6 or 13
	MasterID        *int64
	Quality         *string
	InsertDate      time.Time
	SheetPathDetID  *int64
}
