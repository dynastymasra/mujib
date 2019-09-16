package controller

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/mujib/infrastructure/provider/postgres"

	"github.com/dynastymasra/mujib/config"
	"github.com/sirupsen/logrus"

	"github.com/dynastymasra/mujib/infrastructure/web/formatter"
)

func Ping(postgres *postgres.Connector) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log := logrus.WithField(config.RequestID, r.Context().Value(config.HeaderRequestID))

		if err := postgres.Ping(); err != nil {
			log.WithError(err).Errorln("Failed ping postgres")

			fmt.Fprintf(w, formatter.FailResponse(w, http.StatusInternalServerError, err.Error()).Stringify())
			return
		}

		fmt.Fprint(w, formatter.SuccessResponse(w, http.StatusOK, nil).Stringify())
	}
}
