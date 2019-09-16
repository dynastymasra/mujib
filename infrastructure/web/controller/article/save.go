package article

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dynastymasra/mujib/config"
	"github.com/dynastymasra/mujib/domain"
	"github.com/dynastymasra/mujib/infrastructure/web/formatter"
	"github.com/dynastymasra/mujib/service"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

func Save(service service.ArticleServicer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var requestBody domain.Article

		log := logrus.WithFields(logrus.Fields{
			config.RequestID: r.Context().Value(config.HeaderRequestID),
		})

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.WithError(err).Errorln("Unable to read request body")

			statusCode := http.StatusBadRequest
			w.WriteHeader(statusCode)
			fmt.Fprint(w, formatter.FailResponse(statusCode, err.Error()).Stringify())
			return
		}

		if err := json.Unmarshal(body, &requestBody); err != nil {
			log.WithError(err).WithField("body", string(body)).Errorln("Unable to parse request body")

			statusCode := http.StatusBadRequest
			w.WriteHeader(statusCode)
			fmt.Fprint(w, formatter.FailResponse(statusCode, err.Error()).Stringify())
			return
		}

		validate := validator.New()
		if err := validate.Struct(&requestBody); err != nil {
			log.WithError(err).WithField("body", requestBody).Errorln("Failed validate article request")

			statusCode := http.StatusBadRequest
			w.WriteHeader(statusCode)
			fmt.Fprint(w, formatter.FailResponse(statusCode, err.Error()).Stringify())
			return
		}

		article, err := service.CreateArticle(r.Context(), requestBody)
		if err != nil {
			log.WithError(err).WithField("body", requestBody).Errorln("Failed create new article")

			statusCode := http.StatusInternalServerError
			w.WriteHeader(statusCode)
			fmt.Fprint(w, formatter.FailResponse(statusCode, err.Error()).Stringify())
			return
		}

		statusCode := http.StatusCreated
		w.WriteHeader(statusCode)
		fmt.Fprint(w, formatter.SuccessResponse(statusCode, map[string]interface{}{"id": article.ID}).Stringify())
	}
}
