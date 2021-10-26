package start

import (
	"github.com/dinislamdarkhan/simple-wallet/src/app/store"
	wallet "github.com/dinislamdarkhan/simple-wallet/src/wallet/presenter"
)

type service struct {
	wallet wallet.Service
}

func ServiceFactory(store store.RepositoryStore) *service {
	return &service{
		wallet: wallet.New(store),
	}
}
