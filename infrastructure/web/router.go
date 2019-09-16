package web

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/dynastymasra/mujib/infrastructure/web/controller/article"

	"github.com/dynastymasra/mujib/service"

	"github.com/dynastymasra/mujib/infrastructure/web/middleware"

	"github.com/dynastymasra/mujib/infrastructure/web/controller"

	"github.com/dynastymasra/mujib/config"
	"github.com/dynastymasra/mujib/infrastructure/web/formatter"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func Router(db *gorm.DB, service service.ArticleServicer) *mux.Router {
	router := mux.NewRouter().StrictSlash(true).UseEncodedPath()
	commonHandlers := negroni.New(
		middleware.RequestID(),
		middleware.HTTPStatLogger(),
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
		negroni.WrapFunc(controller.Ping(db)),
	)).Methods(http.MethodGet, http.MethodHead)

	// article group
	router.Handle("/articles", commonHandlers.With(
		negroni.WrapFunc(article.Save(service)),
	)).Methods(http.MethodPost)

	router.Handle("/articles/{article_id:[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}}", commonHandlers.With(
		negroni.WrapFunc(article.FindByID(service)),
	)).Methods(http.MethodGet)

	router.Handle("/articles", commonHandlers.With(
		negroni.WrapFunc(article.FindAll(service)),
	)).Methods(http.MethodGet)

	return router
}
