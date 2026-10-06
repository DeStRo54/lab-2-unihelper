package schedule_updater

import "go.uber.org/zap"

type CronJob struct {
	adminService AdminService
	logger       *zap.SugaredLogger
}

func NewCronJob(adminService AdminService, logger *zap.SugaredLogger) *CronJob {
	return &CronJob{
		adminService: adminService,
		logger:       logger,
	}
}
