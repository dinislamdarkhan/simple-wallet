package cassandra

import (
	"context"
	"time"

	"github.com/dinislamdarkhan/simple-wallet/src/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/dinislamdarkhan/simple-wallet/src/wallet/repository"
	"github.com/gocql/gocql"
)

type cassandraRepository struct {
	db *gocql.Session
}

func New(db *gocql.Session) repository.CassandraRepository {
	return &cassandraRepository{
		db: db,
	}
}

func (c *cassandraRepository) UpdateWalletAmount(ctx context.Context, currency, userID string, amount float64, updatedTime time.Time) error {
	if err := c.db.Query(`UPDATE wallet SET amount = ?, updated_time = ? WHERE currency = ? AND user_id = ?`, amount, updatedTime, currency, userID).Exec(); err != nil {
		return err
	}
	return nil
}

func (c *cassandraRepository) GetWalletAmount(ctx context.Context, currency, userID string) (float64, error) {
	var amount float64
	if err := c.db.Query(`SELECT total_amount FROM wallet WHERE currency = ? AND user_id = ?`, currency, userID).Scan(&amount); err != nil {
		return 0, err
	}
	return amount, nil
}

func (c *cassandraRepository) CheckAmountExists(ctx context.Context, currency, userID string) (count int, err error) {
	logger := logrus.WithContext(ctx)

	if err = c.db.Query(`SELECT count(*) FROM wallet WHERE currency = ? AND user_id = ? LIMIT 1`, currency, userID).Scan(&count); err != nil {
		logger.Error("CheckAmountIfExistsFromCassandra: ", err)
		return 0, errors.CassandraReadError.SetMessage(err.Error())
	}

	return count, nil
}
