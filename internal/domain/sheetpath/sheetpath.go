package sheetpath

import "time"

type SheetPath struct {
	id          int64
	pathID      int64
	matOrderID  int64
	inventoryID *int64
	productID   *int64
	netArea     float64
	usedArea    float64
	used        float64
	insertDate  time.Time
}

func (s SheetPath) ID() int64 {
	return s.id
}

func (s SheetPath) PathID() int64 {
	return s.pathID
}

func (s SheetPath) MatOrderID() int64 {
	return s.matOrderID
}

func (s SheetPath) InventoryID() *int64 {
	return s.inventoryID
}

func (s SheetPath) ProductID() *int64 {
	return s.productID
}

func (s SheetPath) NetArea() float64 {
	return s.netArea
}

func (s SheetPath) UsedArea() float64 {
	return s.usedArea
}

func (s SheetPath) Used() float64 {
	return s.used
}

func (s SheetPath) InsertDate() time.Time {
	return s.insertDate
}

func New(
	id int64,
	pathID int64,
	matOrderID int64,
	inventoryID *int64,
	productID *int64,
	netArea float64,
	usedArea float64,
	used float64,
	insertDate time.Time,
) *SheetPath {
	return &SheetPath{
		id:          id,
		pathID:      pathID,
		matOrderID:  matOrderID,
		inventoryID: inventoryID,
		productID:   productID,
		netArea:     netArea,
		usedArea:    usedArea,
		used:        used,
		insertDate:  insertDate,
	}
}
