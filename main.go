package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli"

	"github.com/dynastymasra/mujib/infrastructure/web"

	"gopkg.in/tylerb/graceful.v1"

	"github.com/dynastymasra/mujib/config"
	"github.com/sirupsen/logrus"
)

func init() {
	config.Load()
	config.Logger().Setup()
}

func main() {
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	clientApp := cli.NewApp()
	clientApp.Name = config.ServiceName
	clientApp.Version = config.Version

	clientApp.Action = func(c *cli.Context) error {
		start(stop)
		return nil
	}

	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}
}

func start(stop chan os.Signal) {
	server := &graceful.Server{
		Timeout: 0,
	}
	go web.Run(server)
	select {
	case sig := <-stop:
		<-server.StopChan()
		logrus.Warningln(fmt.Sprintf("Service shutdown because %+v", sig))

		//if err := provider.Postgres.Close(); err != nil {
		//	logrus.WithError(err).Errorln("Unable to turn off Postgres connections")
		//}

		logrus.Infoln("Postgres Connection closed")
		os.Exit(0)
	}
}
