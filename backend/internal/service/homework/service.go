package homework

import "homewormanager/internal/tx_manager"

type Service struct {
	homeworkRepo          HomeworkRepo
	homeworkStatusService HomeworkStatusService
	manager               *tx_manager.TxManager
}

func NewHomeworkService(homeworkRepo HomeworkRepo, homeworkStatusService HomeworkStatusService, manager *tx_manager.TxManager) *Service {
	return &Service{
		homeworkRepo:          homeworkRepo,
		homeworkStatusService: homeworkStatusService,
		manager:               manager,
	}
}
