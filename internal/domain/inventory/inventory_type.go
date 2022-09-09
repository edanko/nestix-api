package inventory

type InventoryType int64

const (
	UnknownType InventoryType = 0
	PlateType   InventoryType = 1
	PartType    InventoryType = 9
)

func NewProductTypeFromInt64(i int64) InventoryType {
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

func (c InventoryType) Int64() int64 {
	return int64(c)
}

func (c InventoryType) String() string {
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
