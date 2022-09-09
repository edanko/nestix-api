package machine

type Machine struct {
	id   int64
	name string
}

func (m Machine) ID() int64 {
	return m.id
}

func (m Machine) Name() string {
	return m.name
}

func New(
	id int64,
	name string,
) (*Machine, error) {
	if name == "" {
		return nil, ErrValidationName
	}

	return &Machine{
		id:   id,
		name: name,
	}, nil
}
