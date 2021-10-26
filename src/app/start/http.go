package start

import (
	"context"
	"net/http"
	"time"

	"github.com/dinislamdarkhan/simple-wallet/src/pkg/grace"
	walletHTTP "github.com/dinislamdarkhan/simple-wallet/src/wallet/presenter/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func HTTP(services *service, httpAddr string, errs chan<- error) grace.Service {
	mux := http.NewServeMux()

	http.Handle("/v1/", func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				return
			}

			h.ServeHTTP(w, r)
		})
	}(mux))

	mux.Handle("/v1/wallet/", walletHTTP.MakeHandler(services.wallet))
	mux.Handle("/metrics", promhttp.Handler())

	logrus.Infof("start http server on %s", httpAddr)
	logrus.WithFields(logrus.Fields{"transport": "http", "address": httpAddr}).Info("listening simple-wallet")

	server := &http.Server{Addr: httpAddr, Handler: mux}
	go func() {
		errs <- server.ListenAndServe()
	}()

	return grace.NewService("http", func() {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logrus.Errorf("failed to shutdown http server: %v", err)
		}
	})
}
