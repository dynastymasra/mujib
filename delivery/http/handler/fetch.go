package handler

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/dynastymasra/mujib/config"
	"github.com/dynastymasra/mujib/delivery/http/formatter"
	"github.com/dynastymasra/mujib/domain"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func ProductFindByID(service domain.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		v := mux.Vars(r)
		id := v["product_id"]

		log := logrus.WithFields(logrus.Fields{
			config.HeaderRequestID: r.Context().Value(config.HeaderRequestID),
			"product_id":           id,
		})

		product, err := service.FindByID(r.Context(), id)
		if err == gorm.ErrRecordNotFound {
			log.WithError(err).Errorln("Failed get product by id")

			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		if err != nil {
			log.WithError(err).Errorln("Failed get product by id")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, formatter.ObjectResponse(product).Stringify())
	}
}

func ProductFindAll(service domain.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		from := formatter.BuildPaginationFrom(r.FormValue("from"))
		size := formatter.BuildPaginationSize(r.FormValue("size"))

		log := logrus.WithFields(logrus.Fields{
			config.HeaderRequestID: r.Context().Value(config.HeaderRequestID),
			"from":                 from,
			"size":                 size,
		})

		products, err := service.Fetch(r.Context(), from, size)
		if err != nil {
			log.WithError(err).Errorln("Failed fetch product")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, formatter.ObjectResponse(products).Stringify())
	}
}
