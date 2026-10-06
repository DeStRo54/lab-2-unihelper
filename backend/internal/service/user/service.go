package user

import (
	"homewormanager/internal/config"
)

type Service struct {
	userRepo      UserRepo
	userRedisRepo UserRedisRepo
	groupService  GroupService
	config        *config.Config
}

func NewUserService(userRepo UserRepo, userRedisRepo UserRedisRepo, groupService GroupService, config *config.Config) *Service {
	return &Service{
		userRepo:      userRepo,
		config:        config,
		groupService:  groupService,
		userRedisRepo: userRedisRepo,
	}
}
