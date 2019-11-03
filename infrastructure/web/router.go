package web

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/mujib/delivery/http/formatter"
	productHandler "github.com/dynastymasra/mujib/delivery/http/handler"
	"github.com/dynastymasra/mujib/domain"

	"github.com/jinzhu/gorm"

	"github.com/dynastymasra/mujib/infrastructure/web/middleware"

	"github.com/dynastymasra/mujib/infrastructure/web/handler"

	"github.com/dynastymasra/mujib/config"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func Router(db *gorm.DB, service domain.ProductService) *mux.Router {
	router := mux.NewRouter().StrictSlash(true).UseEncodedPath()
	subRouter := router.PathPrefix("/v1/").Subrouter().UseEncodedPath()
	commonHandlers := negroni.New(
		middleware.RequestID(),
	)

	subRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, formatter.FailResponse(config.ErrDataNotFound.Error()).Stringify())
	})

	subRouter.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, formatter.FailResponse(config.ErrDataNotFound.Error()).Stringify())
	})

	// Probes
	subRouter.Handle("/ping", commonHandlers.With(
		negroni.WrapFunc(handler.Ping(db)),
	)).Methods(http.MethodGet, http.MethodHead)

	secretKey := config.SecretKey()

	// product group
	subRouter.Handle("/products", commonHandlers.With(
		middleware.Authorization(secretKey),
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(productHandler.ProductCreate(service)),
	)).Methods(http.MethodPost)

	subRouter.Handle("/products/{product_id}", commonHandlers.With(
		middleware.Authorization(secretKey),
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(productHandler.ProductFindByID(service)),
	)).Methods(http.MethodGet)

	subRouter.Handle("/products", commonHandlers.With(
		middleware.Authorization(secretKey),
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(productHandler.ProductFindAll(service)),
	)).Methods(http.MethodGet)

	subRouter.Handle("/products/{product_id}", commonHandlers.With(
		middleware.Authorization(secretKey),
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(productHandler.ProductUpdate(service)),
	)).Methods(http.MethodPut)

	subRouter.Handle("/products/{product_id}", commonHandlers.With(
		middleware.Authorization(secretKey),
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(productHandler.ProductDelete(service)),
	)).Methods(http.MethodDelete)

	return subRouter
}
