package web

import (
	"net/http"

	"github.com/dynastymasra/mujib/product"

	"github.com/jinzhu/gorm"

	"github.com/dynastymasra/mujib/config"
	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
	"gopkg.in/tylerb/graceful.v1"
)

func Run(server *graceful.Server, db *gorm.DB, service product.Service) {
	logrus.Infoln("Start run web application")

	muxRouter := Router(db, service)

	server.Server = &http.Server{
		Addr: config.ServerAddress(),
		Handler: handlers.RecoveryHandler(
			handlers.PrintRecoveryStack(true),
			handlers.RecoveryLogger(logrus.StandardLogger()),
		)(muxRouter),
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Fatalln("Failed to start server")
	}
}
