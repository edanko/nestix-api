package pagination_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/edanko/nestix-api/pkg/pagination"
)

func TestPagination(t *testing.T) {
	tests := []struct {
		name     string
		wantTime time.Time
		wantID   uuid.UUID
		wantErr  string
	}{
		{
			name:     "successfully decode encoded cursor",
			wantTime: time.Date(2022, 11, 17, 20, 34, 58, 651387237, time.UTC),
			wantID:   uuid.New(),
			wantErr:  "",
		},
		{
			name:     "error with wrong time",
			wantTime: time.Time{},
			wantID:   uuid.New(),
			wantErr:  "pagination.DecodeCursor: invalid time",
		},
		{
			name:     "error with wrong id",
			wantTime: time.Date(2022, 11, 17, 20, 34, 58, 651387237, time.UTC),
			wantID:   uuid.Nil,
			wantErr:  "pagination.DecodeCursor: invalid id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedCursor := pagination.EncodeCursor(tt.wantTime, tt.wantID)

			gotTime, gotID, err := pagination.DecodeCursor(encodedCursor)

			if err != nil {
				require.EqualError(t, err, tt.wantErr)
				return
			}
			require.Equal(t, tt.wantTime, gotTime)
			require.Equal(t, tt.wantID, gotID)

		})
	}
}
