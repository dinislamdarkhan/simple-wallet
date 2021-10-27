package middleware

import (
	"context"
	"time"

	"github.com/dinislamdarkhan/simple-wallet/src/wallet/presenter"
	"github.com/dinislamdarkhan/simple-wallet/src/wallet/presenter/data"
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

func (l *loggingMiddleware) PostDoTransaction(ctx context.Context, req *data.PostDoTransactionRequest) (response *data.PostDoTransactionResponse, err error) {
	defer func(begin time.Time) {
		logMethod(ctx, l.name, "PostDoTransaction", req, response, begin, err)
	}(time.Now())

	response, err = l.service.PostDoTransaction(ctx, req)

	return
}
