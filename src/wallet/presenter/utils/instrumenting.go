package utils

import (
	"context"
	"time"

	"github.com/dinislamdarkhan/simple-wallet/src/wallet/domain"

	"github.com/dinislamdarkhan/simple-wallet/src/wallet/presenter"
	"github.com/go-kit/kit/metrics"
)

func WithInstrumenting(counter metrics.Counter, latency metrics.Histogram, counterE metrics.Counter, service presenter.Service, name string) presenter.Service {
	return &instrumentingMiddleware{
		name:           name,
		requestCount:   counter,
		requestLatency: latency,
		requestError:   counterE,
		service:        service,
	}
}

type instrumentingMiddleware struct {
	name           string
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	requestError   metrics.Counter
	service        presenter.Service
}

func metricMethod(i *instrumentingMiddleware, method string, begin time.Time, err error) {
	i.requestCount.With("method", method).Add(1)
	if err != nil {
		i.requestError.With("method", method).Add(1)
	}
	i.requestLatency.With("method", method).Observe(time.Since(begin).Seconds())
}

func (i *instrumentingMiddleware) PostDoTransaction(ctx context.Context, req *domain.PostDoTransactionRequest) (resp *domain.PostDoTransactionResponse, err error) {
	defer func(begin time.Time) {
		metricMethod(i, "PostDoTransaction", begin, err)
	}(time.Now())

	resp, err = i.service.PostDoTransaction(ctx, req)

	return
}

func (i *instrumentingMiddleware) GetWalletBalance(ctx context.Context, req *domain.GetWalletBalanceRequest) (resp *domain.GetWalletBalanceResponse, err error) {
	defer func(begin time.Time) {
		metricMethod(i, "GetWalletBalance", begin, err)
	}(time.Now())

	resp, err = i.service.GetWalletBalance(ctx, req)

	return
}
