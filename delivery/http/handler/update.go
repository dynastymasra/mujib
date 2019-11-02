package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dynastymasra/mujib/config"
	"github.com/dynastymasra/mujib/delivery/http/formatter"
	"github.com/dynastymasra/mujib/domain"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

func ProductUpdate(service domain.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		v := mux.Vars(r)
		id := v["product_id"]

		log := logrus.WithFields(logrus.Fields{
			config.RequestID: r.Context().Value(config.HeaderRequestID),
			"product_id":     id,
		})

		var reqBody domain.Product

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.WithError(err).Errorln("Unable to read request body")

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		if err := json.Unmarshal(body, &reqBody); err != nil {
			log.WithError(err).WithField("body", string(body)).Errorln("Unable to parse request body")

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		if err := validator.New().Struct(&reqBody); err != nil {
			log.WithError(err).WithField("body", reqBody).Errorln("Failed validate product request")

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		product, err := service.Update(r.Context(), id, reqBody)
		if err != nil {
			log.WithError(err).WithField("body", reqBody).Errorln("Failed create new product")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, formatter.ObjectResponse(product).Stringify())
	}
}
