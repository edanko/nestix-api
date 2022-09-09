package adapters

import "github.com/rs/zerolog/log"

var (
	logf = func(tenantID string, query string, args ...any) {
		log.Debug().
			Str("tenant_id", tenantID).
			Str("query", query).
			Interface("args", args).
			Msg("db")
	}

	logerror = func(err error) error {
		log.Error().
			Err(err).
			Msg("db")

		return err
	}
)
