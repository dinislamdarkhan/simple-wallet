package cassandra

import (
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
