package schedule_updater

import (
	"context"
	"go.uber.org/zap"
)

func (j *CronJob) Run() {
	err := j.adminService.UpdateClasses(context.Background())
	if err != nil {
		j.logger.Errorw("Unhandled error occurred while updating classes",
			zap.Error(err),
		)
		return
	}
	j.logger.Info("Job completed successfully")
}
