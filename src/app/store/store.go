package store

import (
	walletRepository "github.com/dinislamdarkhan/simple-wallet/src/wallet/repository"
	walletCassandra "github.com/dinislamdarkhan/simple-wallet/src/wallet/repository/cassandra"
	"github.com/gocql/gocql"
)

type RepositoryStore interface {
	WalletCassandra() walletRepository.CassandraRepository
}

func New(cass *gocql.Session) RepositoryStore {
	return &repositoryStore{
		nomenclatureCassandra: walletCassandra.New(cass),
	}
}

type repositoryStore struct {
	nomenclatureCassandra walletRepository.CassandraRepository
}

func (r *repositoryStore) WalletCassandra() walletRepository.CassandraRepository {
	return r.nomenclatureCassandra
}
