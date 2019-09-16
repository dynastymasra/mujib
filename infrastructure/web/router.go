package web

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/mujib/infrastructure/provider/postgres"

	"github.com/dynastymasra/mujib/infrastructure/web/controller"

	"github.com/dynastymasra/mujib/config"
	"github.com/dynastymasra/mujib/infrastructure/web/formatter"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func Router(postgres *postgres.Connector) *mux.Router {
	router := mux.NewRouter().StrictSlash(true).UseEncodedPath()
	commonHandlers := negroni.New(
	//middleware.HTTPStatLogger(),
	//middleware.RequestID(),
	)

	subRouter := router.PathPrefix("/v1/").Subrouter().UseEncodedPath()

	subRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, formatter.FailResponse(w, http.StatusNotFound, config.ErrDataNotFound.Error()).Stringify())
	})

	// Probes
	subRouter.Handle("/ping", commonHandlers.With(
		negroni.WrapFunc(controller.Ping(postgres)),
	)).Methods(http.MethodGet, http.MethodHead)

	return router
}
