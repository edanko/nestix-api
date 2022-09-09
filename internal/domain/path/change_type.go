package path

type ChangeType int64

const (
	UnChanged               ChangeType = 0
	MaterialChanged         ChangeType = 1
	GeometryChanged         ChangeType = 2
	SomeWrongValue          ChangeType = 3
	BevelTechnologyChanged  ChangeType = 4
	OuterProfileChanged     ChangeType = 8
	TechnologyChanged       ChangeType = 16
	TextOrAnnotationChanged ChangeType = 32
)

func NewChangeTypeFromInt64(i *int64) ChangeType {
	if i == nil {
		return UnChanged
	}
	switch *i {
	case 0:
		return UnChanged
	case 1:
		return MaterialChanged
	case 2:
		return GeometryChanged
	case 3:
		return SomeWrongValue
	case 4:
		return BevelTechnologyChanged
	case 8:
		return OuterProfileChanged
	case 16:
		return TechnologyChanged
	case 32:
		return TextOrAnnotationChanged
	}
	return ChangeType(*i)
}

func (c ChangeType) Int64() *int64 {
	if c == UnChanged {
		return nil
	}
	val := int64(c)
	return &val
}

func (c ChangeType) String() string {
	switch c {
	case UnChanged:
		return "unchanged"
	case MaterialChanged:
		return "material_changed"
	case GeometryChanged:
		return "geometry_changed"
	case SomeWrongValue:
		return "some_wrong_value"
	case BevelTechnologyChanged:
		return "bevel_technology_changed"
	case OuterProfileChanged:
		return "outer_profile_changed"
	case TechnologyChanged:
		return "technology_changed"

	case TextOrAnnotationChanged:
		return "text_or_annotation_changed"
	}
	return "unknown"
}
