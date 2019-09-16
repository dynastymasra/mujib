package article

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/mujib/infrastructure/web/formatter"
	"github.com/jinzhu/gorm"

	"github.com/dynastymasra/mujib/config"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/dynastymasra/mujib/service"
)

func FindByID(service service.ArticleServicer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		v := mux.Vars(r)
		id := v["article_id"]

		log := logrus.WithFields(logrus.Fields{
			config.HeaderRequestID: r.Context().Value(config.HeaderRequestID),
			"article_id":           id,
		})

		article, err := service.FindArticleByID(r.Context(), id)
		if err == gorm.ErrRecordNotFound {
			log.WithError(err).Warn("Article data not found")

			statusCode := http.StatusNotFound
			w.WriteHeader(statusCode)
			fmt.Fprint(w, formatter.FailResponse(statusCode, err.Error()).Stringify())
			return
		}
		if err != nil {
			log.WithError(err).Errorln("Error get article by id from db")

			statusCode := http.StatusInternalServerError
			w.WriteHeader(statusCode)
			fmt.Fprint(w, formatter.FailResponse(statusCode, err.Error()).Stringify())
			return
		}

		response := []interface{}{article}
		statusCode := http.StatusOK
		w.WriteHeader(statusCode)
		fmt.Fprint(w, formatter.SuccessResponse(statusCode, response).Stringify())
	}
}
