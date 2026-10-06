package middleware

type MwManager struct {
	userRedisRepo UserRedisRepo
	userRepo      UserRepo
}

func NewMwManager(userRedisRepo UserRedisRepo, userRepo UserRepo) *MwManager {
	return &MwManager{
		userRedisRepo: userRedisRepo,
		userRepo:      userRepo,
	}
}
