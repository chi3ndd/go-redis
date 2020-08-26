package redis

import (
	"os"
	"time"

	r "github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
)

type (
	Connector struct {
		Addr     string
		Password string
		Database int
		client   *r.Client
		Logger   *logrus.Logger
	}
)

func (con *Connector) Initiation() error {
	// Initiation logger
	con.Logger = &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: time.RFC3339,
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}
	// Initiation Redis Client
	con.client = r.NewClient(&r.Options{
		Addr:       con.Addr,
		Password:   con.Password,
		DB:         con.Database,
		MaxRetries: 5,
	})
	con.Logger.Infof("Initializing connection to Redis server [%s] - database (%d)", con.Addr, con.Database)
	// Success
	return con.client.Ping().Err()
}
