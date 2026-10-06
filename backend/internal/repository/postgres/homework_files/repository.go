package homework_files

import "homewormanager/internal/tx_manager"

type Repo struct {
	manager *tx_manager.TxManager
}

func NewHomeworkFilesRepo(manager *tx_manager.TxManager) *Repo {
	return &Repo{manager: manager}
}
