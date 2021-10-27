package start

import (
	"github.com/dinislamdarkhan/simple-wallet/src/app/store"
	"github.com/dinislamdarkhan/simple-wallet/src/wallet"
	walletPresenter "github.com/dinislamdarkhan/simple-wallet/src/wallet/presenter"
)

type service struct {
	wallet walletPresenter.Service
}

func ServiceFactory(store store.RepositoryStore) *service {
	return &service{
		wallet: wallet.New(store),
	}
}
