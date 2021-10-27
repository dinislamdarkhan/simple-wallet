package wallet

import (
	"github.com/dinislamdarkhan/simple-wallet/src/app/store"
	"github.com/dinislamdarkhan/simple-wallet/src/wallet/presenter"
	"github.com/dinislamdarkhan/simple-wallet/src/wallet/presenter/utils"
	kitmetrics "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	serviceName = "simple-wallet"
)

func New(store store.RepositoryStore) (service presenter.Service) {
	service = presenter.New(store)
	service = utils.WithLogging(service, serviceName)
	fieldKeys := []string{"method"}
	service = utils.WithInstrumenting(
		kitmetrics.NewCounterFrom(prometheus.CounterOpts{
			Namespace: "simple-wallet-API",
			Subsystem: "wallet_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitmetrics.NewSummaryFrom(prometheus.SummaryOpts{
			Namespace: "simple-wallet-API",
			Subsystem: "wallet_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		kitmetrics.NewCounterFrom(prometheus.CounterOpts{
			Namespace: "simple-wallet-API",
			Subsystem: "wallet_service",
			Name:      "error_count",
			Help:      "Number of error requests received.",
		}, fieldKeys),
		service,
		serviceName,
	)

	return
}
