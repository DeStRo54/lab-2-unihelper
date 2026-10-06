package class

import (
	"homewormanager/internal/tx_manager"
)

type Service struct {
	groupService   GroupService
	classRepo      ClassRepo
	subjectService SubjectService
	manager        *tx_manager.TxManager
}

func NewClassService(groupService GroupService, classRepo ClassRepo, subjectService SubjectService, manager *tx_manager.TxManager) *Service {
	return &Service{
		groupService:   groupService,
		classRepo:      classRepo,
		subjectService: subjectService,
		manager:        manager,
	}
}
