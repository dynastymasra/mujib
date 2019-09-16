package postgres

import (
	"errors"

	"github.com/dynastymasra/mujib/config"
	"github.com/jinzhu/gorm"
	"github.com/matryer/resync"
	"github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	Client *connector
	db     *gorm.DB
	err    error
	once   resync.Once
)

type connector struct {
	DB *gorm.DB
}

func Connect(config config.PostgresConfig) (*connector, error) {
	dbURL := config.ConnectionString()

	once.Do(func() {
		db, err = gorm.Open("postgres", dbURL)
		if err != nil {
			logrus.WithError(err).WithField("db_url", dbURL).Errorln("Cannot connect to DB")
			return
		}

		db.DB().SetMaxIdleConns(config.MaxIdleConn())
		db.DB().SetMaxOpenConns(config.MaxOpenConn())

		if err := db.DB().Ping(); err != nil {
			logrus.WithError(err).Errorln("Cannot ping database")
			return
		}

		db.LogMode(config.LogEnabled())

		Client = &connector{DB: db}
	})

	return Client, err
}

func (c *connector) Ping() error {
	if c.DB == nil {
		return errors.New("does't have database data")
	}
	return c.DB.DB().Ping()
}

func (c *connector) Close() error {
	if c.DB == nil {
		return errors.New("does't have database data")
	}
	return c.DB.DB().Close()
}

func (c *connector) Reset() {
	once.Reset()
}
