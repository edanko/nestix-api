package order

type RowType int64

const (
	Unknown  RowType = 0
	Part     RowType = 1
	Material RowType = 10
)

func NewRowTypeFromInt64(i int64) RowType {
	switch i {
	case 1:
		return Part
	case 2:
		return Material
	}
	return Unknown
}

func (c RowType) Int64() int64 {
	return int64(c)
}

func (c RowType) String() string {
	switch c {
	case Part:
		return "part"
	case Material:
		return "material"
	}
	return "unknown"
}
