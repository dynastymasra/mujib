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
	db   *gorm.DB
	err  error
	once resync.Once
)

type Connector struct {
	DB *gorm.DB
}

func Connect(config config.PostgresConfig) (*Connector, error) {
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
	})

	return &Connector{DB: db}, err
}

func (c *Connector) Ping() error {
	if c.DB == nil {
		return errors.New("does't have database data")
	}
	return c.DB.DB().Ping()
}

func (c *Connector) Close() error {
	if c.DB == nil {
		return errors.New("does't have database data")
	}
	return c.DB.DB().Close()
}

func (c *Connector) Reset() {
	once.Reset()
}
