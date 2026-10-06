package user

import "homewormanager/internal/tx_manager"

type Repo struct {
	manager *tx_manager.TxManager
}

func NewUserRepo(manager *tx_manager.TxManager) *Repo {
	return &Repo{manager: manager}
}
