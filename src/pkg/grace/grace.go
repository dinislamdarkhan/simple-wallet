package grace

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

type Graceful struct {
	fifo []Service
}

func KillSoftly(fifo ...Service) Graceful {
	return Graceful{fifo: fifo}
}

func (g Graceful) Shutdown(errs chan error, conn Connections) {
	exitCode := 1

	go func() {
		c := make(chan os.Signal, 3)

		signal.Notify(c, syscall.SIGTERM)
		signal.Notify(c, syscall.SIGINT)
		signal.Notify(c, syscall.SIGILL)

		exitCode = 0

		errs <- fmt.Errorf("%s", <-c)
	}()

	err := <-errs
	logrus.Printf("terminated: %v\n", err)

	for _, s := range g.fifo {
		if s != nil {
			serviceName := s.Name()
			logrus.Printf("shutdown service %s started\n", serviceName)
			s.Stop()
			logrus.Printf("shutdown service %s completed\n", serviceName)
		}
	}

	conn.Close()

	logrus.Println("Connections closed")

	os.Exit(exitCode)
}
