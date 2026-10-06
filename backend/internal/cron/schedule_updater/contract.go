package schedule_updater

import "context"

type AdminService interface {
	UpdateClasses(ctx context.Context) error
}
