package sheetpathdet

type DetailType int64

const (
	Unknown DetailType = 0
	Part    DetailType = 1
	Remnant DetailType = 2
)

func NewDetailTypeFromInt64(i int64) DetailType {
	switch i {
	case 1:
		return Part
	case 2:
		return Remnant
	}
	return Unknown
}

func (c DetailType) Int64() int64 {
	return int64(c)
}

func (c DetailType) String() string {
	switch c {
	case Part:
		return "part"
	case Remnant:
		return "remnant"
	}
	return "unknown"
}
