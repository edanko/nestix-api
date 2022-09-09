package handlers

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"

	grpcutils "github.com/edanko/nestix-api/pkg/grpc"

	"google.golang.org/grpc"
)

func BuildReadinessHandler(conn *grpc.ClientConn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !grpcutils.IsConnectionServing(r.Context(), "kind", conn) {
			log.Warn().Msg(fmt.Sprintf("gRPC connection %s is not serving", "kind"))
			// return apperrors.New(fmt.Sprintf("gRPC connection %s is not serving", name))
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// // BuildReadinessHandler provides readiness handler
// func BuildReadinessHandler(sqlConn *sql.DB, mongoConn *mongo.Client, connMap map[string]*grpc.ClientConn) http.Handler {
// 	fn := func(w http.ResponseWriter, r *http.Request) error {
// 		if sqlConn != nil {
// 			if err := sqlConn.PingContext(r.Context()); err != nil {
// 				return apperrors.Wrap(err)
// 			}
// 		}
// 		if mongoConn != nil {
// 			if err := mongoConn.Ping(r.Context(), nil); err != nil {
// 				return apperrors.Wrap(err)
// 			}
// 		}

// 		for name, conn := range connMap {
// 			if !grpcutils.IsConnectionServing(r.Context(), name, conn) {
// 				return apperrors.New(fmt.Sprintf("gRPC connection %s is not serving", name))
// 			}
// 		}

// 		w.WriteHeader(http.StatusNoContent)

// 		return nil
// 	}

// 	return httpjson.HandlerFunc(fn)
// }
