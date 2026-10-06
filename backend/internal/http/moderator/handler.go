package moderator

import (
	"github.com/gofiber/fiber/v2"
	"homewormanager/internal/entity"

	"homewormanager/internal/middleware"
)

type Handler struct {
	moderatorService ModeratorService
}

func NewModeratorHandler(moderatorService ModeratorService) *Handler {
	return &Handler{
		moderatorService: moderatorService,
	}
}

func MapModeratorRoutes(g fiber.Router, h *Handler, mw *middleware.MwManager) {
	g.Post("/addHomework/class", mw.Auth(), mw.AllowedRoles([]entity.Role{entity.RoleGlobalAdmin, entity.RoleGroupModerator}), h.AddHomeworkToClass())
	g.Post("/addHomework/date", mw.Auth(), mw.AllowedRoles([]entity.Role{entity.RoleGlobalAdmin, entity.RoleGroupModerator}), h.AddHomeworkToDate())
	g.Delete("/delete", mw.Auth(), mw.AllowedRoles([]entity.Role{entity.RoleGlobalAdmin, entity.RoleGroupModerator}), h.DeleteHomework())
	g.Patch("/update", mw.Auth(), mw.AllowedRoles([]entity.Role{entity.RoleGlobalAdmin, entity.RoleGroupModerator}), h.UpdateHomework())
	g.Post("/note/add/:subjectID", mw.Auth(), mw.AllowedRoles([]entity.Role{entity.RoleGlobalAdmin, entity.RoleGroupModerator}), h.AddNote())
	g.Delete("/note/delete/:noteID", mw.Auth(), mw.AllowedRoles([]entity.Role{entity.RoleGlobalAdmin, entity.RoleGroupModerator}), h.DeleteNote())
	g.Patch("/note/update/:noteID", mw.Auth(), mw.AllowedRoles([]entity.Role{entity.RoleGlobalAdmin, entity.RoleGroupModerator}), h.UpdateNote())
	g.Patch("/key/regenerate/:groupID", mw.Auth(), mw.AllowedRoles([]entity.Role{entity.RoleGlobalAdmin, entity.RoleGroupModerator}), h.RegenerateGroupRegisterKey())
	g.Get("/key/get", mw.Auth(), mw.AllowedRoles([]entity.Role{entity.RoleGlobalAdmin, entity.RoleGroupModerator}), h.GetGroupRegisterKey())
	g.Post("/addFile/:homeworkID", mw.Auth(), mw.AllowedRoles([]entity.Role{entity.RoleGlobalAdmin, entity.RoleGroupModerator}), h.AddFileToHomework())
	g.Delete("/file/:fileID", mw.Auth(), mw.AllowedRoles([]entity.Role{entity.RoleGlobalAdmin, entity.RoleGroupModerator}), h.DeleteFileFromHomework())
}
