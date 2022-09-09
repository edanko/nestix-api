package sheetpathdet

type Remnant struct {
	id          int64
	sheetPathID int64
	productID   int64
	detailCount int64
	area        float64
	sequenceNo  int64
}

func (s Remnant) ID() int64 {
	return s.id
}

func (s Remnant) SheetPathID() int64 {
	return s.sheetPathID
}

func (s Remnant) ProductID() int64 {
	return s.productID
}

func (s Remnant) DetailCount() int64 {
	return s.detailCount
}

func (s Remnant) Area() float64 {
	return s.area
}

func (s Remnant) SequenceNo() int64 {
	return s.sequenceNo
}

func NewRemnant(
	id int64,
	sheetPathID int64,
	productID *int64,
	detailCount int64,
	area float64,
	sequenceNo int64,
) (*Remnant, error) {
	if sequenceNo < 1 {
		return nil, ErrValidationSequenceNo
	}

	if productID == nil {
		return nil, ErrValidationRemnantProductID
	}

	return &Remnant{
		id:          id,
		sheetPathID: sheetPathID,
		productID:   *productID,
		detailCount: detailCount,
		area:        area,
		sequenceNo:  sequenceNo,
	}, nil
}
