package product

type ProductType int64

const (
	UnknownType ProductType = 0
	PlateType   ProductType = 1
	PartType    ProductType = 9
	RemnantType ProductType = 12
	FrameType   ProductType = 15
)

func NewProductTypeFromInt64(i int64) ProductType {
	switch i {
	case 1:
		return PlateType
	case 9:
		return PartType
	case 12:
		return RemnantType
	case 15:
		return FrameType
	}
	return UnknownType
}

func (c ProductType) Int64() int64 {
	return int64(c)
}

func (c ProductType) String() string {
	switch c {
	case PlateType:
		return "plate"
	case PartType:
		return "part"
	case RemnantType:
		return "remnant"
	case FrameType:
		return "frame"
	}
	return "unknown"
}
