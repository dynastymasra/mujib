package controller

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/mujib/infrastructure/database/postgres"
	"github.com/jinzhu/gorm"

	"github.com/dynastymasra/mujib/config"
	"github.com/sirupsen/logrus"

	"github.com/dynastymasra/mujib/infrastructure/web/formatter"
)

func Ping(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log := logrus.WithField(config.RequestID, r.Context().Value(config.HeaderRequestID))

		if err := postgres.Ping(db); err != nil {
			log.WithError(err).Errorln("Failed ping postgres")

			statusCode := http.StatusInternalServerError
			w.WriteHeader(statusCode)
			fmt.Fprintf(w, formatter.FailResponse(statusCode, err.Error()).Stringify())
			return
		}

		statusCode := http.StatusOK
		w.WriteHeader(statusCode)
		fmt.Fprint(w, formatter.SuccessResponse(statusCode, nil).Stringify())
	}
}
