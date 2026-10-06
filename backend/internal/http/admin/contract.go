package admin

import (
	"context"
	"homewormanager/internal/entity"

	adminService "homewormanager/internal/service/admin"
)

type AdminService interface {
	AddGroup(ctx context.Context, req adminService.AddGroup) error
	UpdateClasses(ctx context.Context) error
	RefreshAllData(ctx context.Context) error
	UpdateRole(ctx context.Context, req adminService.UpdateUserRole) error
	GetAllUsers(ctx context.Context) ([]entity.UserFullInfo, error)
	GetAllGroups(ctx context.Context) ([]entity.Group, error)
	RegenerateAllRegisterKeys(ctx context.Context) error
}
