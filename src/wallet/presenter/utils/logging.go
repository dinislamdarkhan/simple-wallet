package utils

import (
	"context"
	"time"

	"github.com/dinislamdarkhan/simple-wallet/src/wallet/domain"

	"github.com/dinislamdarkhan/simple-wallet/src/wallet/presenter"
	"github.com/sirupsen/logrus"
)

func WithLogging(service presenter.Service, name string) presenter.Service {
	return &loggingMiddleware{
		name:    name,
		service: service,
	}
}

type loggingMiddleware struct {
	name    string
	service presenter.Service
}

func logMethod(ctx context.Context, name, method string, request, response interface{}, begin time.Time, err error) {
	log := logrus.WithContext(ctx).
		WithFields(logrus.Fields{
			"name":     name,
			"method":   method,
			"request":  request,
			"response": response,
			"begin":    begin,
		})

	if err != nil {
		log.WithError(err).Error()
	} else {
		log.Info()
	}
}

func (l *loggingMiddleware) PostDoTransaction(ctx context.Context, req *domain.PostDoTransactionRequest) (resp *domain.PostDoTransactionResponse, err error) {
	defer func(begin time.Time) {
		logMethod(ctx, l.name, "PostDoTransaction", req, resp, begin, err)
	}(time.Now())

	resp, err = l.service.PostDoTransaction(ctx, req)

	return
}

func (l *loggingMiddleware) GetWalletBalance(ctx context.Context, req *domain.GetWalletBalanceRequest) (resp *domain.GetWalletBalanceResponse, err error) {
	defer func(begin time.Time) {
		logMethod(ctx, l.name, "GetWalletBalance", req, resp, begin, err)
	}(time.Now())

	resp, err = l.service.GetWalletBalance(ctx, req)

	return
}
