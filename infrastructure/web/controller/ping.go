package controller

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/mujib/infrastructure/web/formatter"
)

func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		//log := logrus.WithField(config.RequestID, r.Context().Value(config.HeaderRequestID))

		fmt.Fprint(w, formatter.SuccessResponse(w, http.StatusOK, nil).Stringify())
	}
}
