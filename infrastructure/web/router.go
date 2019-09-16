package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true).UseEncodedPath()
	negroni.New(
	//middleware.HTTPStatLogger(),
	//middleware.RequestID(),
	)

	subRouter := router.PathPrefix("/v1/").Subrouter().UseEncodedPath()

	subRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		//fmt.Fprintf(w, formatter.FailResponse(config.ErrDataNotFound.Error()).Stringify())
	})

	return router
}
