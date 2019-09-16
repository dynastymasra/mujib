package web

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/mujib/infrastructure/web/controller/article"

	"github.com/dynastymasra/mujib/service"

	"github.com/dynastymasra/mujib/infrastructure/web/middleware"

	"github.com/dynastymasra/mujib/infrastructure/provider/postgres"

	"github.com/dynastymasra/mujib/infrastructure/web/controller"

	"github.com/dynastymasra/mujib/config"
	"github.com/dynastymasra/mujib/infrastructure/web/formatter"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func Router(postgres *postgres.Connector, service service.ArticleServicer) *mux.Router {
	router := mux.NewRouter().StrictSlash(true).UseEncodedPath()
	commonHandlers := negroni.New(
		middleware.HTTPStatLogger(),
		middleware.RequestID(),
	)

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, formatter.FailResponse(http.StatusNotFound, config.ErrDataNotFound.Error()).Stringify())
	})

	router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, formatter.FailResponse(http.StatusMethodNotAllowed, config.ErrDataNotFound.Error()).Stringify())
	})

	// Probes
	router.Handle("/ping", commonHandlers.With(
		negroni.WrapFunc(controller.Ping(postgres)),
	)).Methods(http.MethodGet, http.MethodHead)

	// Order group
	router.Handle("/articles", commonHandlers.With(
		negroni.WrapFunc(article.Save(service)),
	)).Methods(http.MethodPost)

	return router
}
