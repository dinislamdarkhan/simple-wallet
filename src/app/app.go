package app

import (
	"fmt"
	"os"
	"time"

	"github.com/dinislamdarkhan/simple-wallet/src/app/config"
	"github.com/dinislamdarkhan/simple-wallet/src/app/conns"
	"github.com/sirupsen/logrus"
)

func Run(httpAddr string) {
	cfg, err := config.GetConfigs()
	if err != nil {
		panic(fmt.Errorf("get config file error: %s", err))
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.Level(cfg.Logrus.Level))

	errs := make(chan error, 2)
	connections, err := conns.New(cfg.Cassandra)
	if err != nil {
		logrus.Fatalf("create connections error: %v", err)
	}
}
