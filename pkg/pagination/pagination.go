package pagination

import (
	"encoding/base64"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func DecodeCursor(encodedCursor string) (time.Time, uuid.UUID, error) {
	if encodedCursor == "" {
		return time.Time{}, uuid.Nil, errors.New("pagination.DecodeCursor: empty cursor")
	}
	byt, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return time.Time{}, uuid.Nil, err
	}

	idBytes := byt[:16]
	timeBytes := byt[16:]

	idParsed := uuid.UUID{}
	err = idParsed.UnmarshalBinary(idBytes)
	if err != nil {
		return time.Time{}, uuid.Nil, err
	}
	if idParsed == uuid.Nil {
		return time.Time{}, uuid.Nil, errors.New("pagination.DecodeCursor: invalid id")
	}

	timeParsed := time.Time{}
	timeParsed.UnmarshalBinary(timeBytes)
	if err != nil {
		return time.Time{}, uuid.Nil, err
	}
	if timeParsed.IsZero() {
		return time.Time{}, uuid.Nil, errors.New("pagination.DecodeCursor: invalid time")
	}

	return timeParsed, idParsed, nil
}

func EncodeCursor(t time.Time, id uuid.UUID) string {
	idBytes, _ := id.MarshalBinary()
	timeBytes, _ := t.MarshalBinary()
	bytes := append(idBytes, timeBytes...)

	return base64.StdEncoding.EncodeToString(bytes)
}
