package sheetpathdet

type Part struct {
	id          int64
	sheetPathID int64
	orderID     int64
	partID      int64
	detailCount int64
	detailCode  string
	usedArea    float64
	area        float64
	slagArea    float64
	nestScrap   float64
	claimArea   float64
	returnArea  float64
	sequenceNo  int64
}

func (s Part) ID() int64 {
	return s.id
}

func (s Part) SheetPathID() int64 {
	return s.sheetPathID
}

func (s Part) OrderID() int64 {
	return s.orderID
}

func (s Part) PartID() int64 {
	return s.partID
}

func (s Part) DetailCount() int64 {
	return s.detailCount
}

func (s Part) DetailCode() string {
	return s.detailCode
}

func (s Part) UsedArea() float64 {
	return s.usedArea
}

func (s Part) Area() float64 {
	return s.area
}

func (s Part) SlagArea() float64 {
	return s.slagArea
}

func (s Part) NestScrap() float64 {
	return s.nestScrap
}

func (s Part) ClaimArea() float64 {
	return s.claimArea
}

func (s Part) ReturnArea() float64 {
	return s.returnArea
}

func (s Part) SequenceNo() int64 {
	return s.sequenceNo
}

func NewPart(
	id int64,
	sheetPathID int64,
	orderID *int64,
	partID *int64,
	detailCount int64,
	detailCode string,
	usedArea float64,
	area float64,
	slagArea float64,
	nestScrap float64,
	claimArea float64,
	returnArea float64,
	sequenceNo int64,
) (*Part, error) {
	if sequenceNo < 1 {
		return nil, ErrValidationSequenceNo
	}

	if orderID == nil {
		return nil, ErrValidationPartOrderID
	}
	if partID == nil {
		return nil, ErrValidationPartPartID
	}

	return &Part{
		id:          id,
		sheetPathID: sheetPathID,
		orderID:     *orderID,
		partID:      *partID,
		detailCount: detailCount,
		detailCode:  detailCode,
		usedArea:    usedArea,
		area:        area,
		slagArea:    slagArea,
		nestScrap:   nestScrap,
		claimArea:   claimArea,
		returnArea:  returnArea,
		sequenceNo:  sequenceNo,
	}, nil
}
