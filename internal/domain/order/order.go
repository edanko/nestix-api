package order

import "time"

type Order struct {
	id         int64   // both
	productID  int64   // both
	partID     *int64  // part
	orderNo    *string // part
	lineNo     int64   // both
	parentName *string // ?
	section    *string // part
	mrpLineNo  *string // part
	count      int64   // both
	// rejected         *int64
	dueDate     *time.Time
	nestedCount *int64 // part
	// rotate           *int64
	// status           *int64
	matStd    *string  // both
	rowType   RowType  // both
	infoTxt   *string  // ?
	totWeight *float64 // mat
	weight    *float64 // both, nullable?
	thick     *float64 // both
	width     *float64 // both
	length    *float64 // both
	purCode   *int64   // both, zero for mat, NULL for part
	pathName  *string  // mat
	// ready            *int64
	// source           *int64
	// type *int64
	// mirror           *int64
	// prodInfo         *string
	// partKey          *int64
	// attributeKey     *string
	// logisticalKey    *int64
	// centerOfGravityX *float64
	// centerOfGravityY *float64
	// centerOfGravityZ *float64
	// projectName      *string
	// assemblySequence *int64
	// name            *string
	// customerID      *int64
	// customerOrderNo *string
	// customerMark    *string
	// workPhases      *string
	// info1           *string
	// info2           *string
	// info3           *string
	// chargeNo        *string
	insertDate time.Time
	created    time.Time
	creator    string
	changed    *time.Time
	changer    *string
	// sourceOrderNo *string
	// partSide      *int64
}

func (o Order) ID() int64 {
	return o.id
}

func (o Order) ProductID() int64 {
	return o.productID
}

func (o Order) PartID() *int64 {
	return o.partID
}

func (o Order) OrderNo() *string {
	return o.orderNo
}

func (o Order) LineNo() int64 {
	return o.lineNo
}

func (o Order) ParentName() *string {
	return o.parentName
}

func (o Order) Section() *string {
	return o.section
}

func (o Order) MrpLineNo() *string {
	return o.mrpLineNo
}

func (o Order) Count() int64 {
	return o.count
}

func (o Order) DueDate() *time.Time {
	return o.dueDate
}

func (o Order) NestedCount() *int64 {
	return o.nestedCount
}

func (o Order) MatStd() *string {
	return o.matStd
}

func (o Order) RowType() RowType {
	return o.rowType
}

func (o Order) InfoTxt() *string {
	return o.infoTxt
}

func (o Order) TotWeight() *float64 {
	return o.totWeight
}

func (o Order) Weight() *float64 {
	return o.weight
}

func (o Order) Thick() *float64 {
	return o.thick
}

func (o Order) Width() *float64 {
	return o.width
}

func (o Order) Length() *float64 {
	return o.length
}

func (o Order) PurCode() *int64 {
	return o.purCode
}

func (o Order) PathName() *string {
	return o.pathName
}

func (o Order) InsertDate() time.Time {
	return o.insertDate
}

func (o Order) Created() time.Time {
	return o.created
}

func (o Order) Creator() string {
	return o.creator
}

func (o Order) Changed() *time.Time {
	return o.changed
}

func (o Order) Changer() *string {
	return o.changer
}

func New(
	id int64,
	productID int64,
	partID *int64,
	orderNo *string,
	lineNo int64,
	parentName *string,
	section *string,
	mrpLineNo *string,
	count int64,
	dueDate *time.Time,
	nestedCount *int64,
	matStd *string,
	rowType RowType,
	infoTxt *string,
	totWeight *float64,
	weight *float64,
	thick *float64,
	width *float64,
	length *float64,
	purCode *int64,
	pathName *string,
	insertDate time.Time,
	creator string,
	created time.Time,
	changer *string,
	changed *time.Time,
) *Order {
	return &Order{
		id:          id,
		productID:   productID,
		partID:      partID,
		orderNo:     orderNo,
		lineNo:      lineNo,
		parentName:  parentName,
		section:     section,
		mrpLineNo:   mrpLineNo,
		count:       count,
		dueDate:     dueDate,
		nestedCount: nestedCount,
		matStd:      matStd,
		rowType:     rowType,
		infoTxt:     infoTxt,
		totWeight:   totWeight,
		weight:      weight,
		thick:       thick,
		width:       width,
		length:      length,
		purCode:     purCode,
		pathName:    pathName,
		insertDate:  insertDate,
		created:     created,
		creator:     creator,
		changed:     changed,
		changer:     changer,
	}
}
