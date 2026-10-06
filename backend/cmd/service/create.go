package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"homewormanager/internal/client/http/s3"
	"homewormanager/internal/config"
	"homewormanager/internal/err_handler"
	adminHandlers "homewormanager/internal/http/admin"
	groupHandlers "homewormanager/internal/http/group"
	homeworkStatusHandlers "homewormanager/internal/http/homework_status"
	moderatorHandlers "homewormanager/internal/http/moderator"
	scheduleHandlers "homewormanager/internal/http/schedule"
	subjectHandlers "homewormanager/internal/http/subject"
	subjectNoteHandlers "homewormanager/internal/http/subjectnote"
	userHandlers "homewormanager/internal/http/user"
	middleware "homewormanager/internal/middleware"
	classRepo "homewormanager/internal/repository/postgres/class"
	groupRepo "homewormanager/internal/repository/postgres/group"
	homeworkRepo "homewormanager/internal/repository/postgres/homework"
	homeworkFilesRepo "homewormanager/internal/repository/postgres/homework_files"
	homeworkStatusRepo "homewormanager/internal/repository/postgres/homework_status"
	subjectRepo "homewormanager/internal/repository/postgres/subject"
	subjectNoteRepo "homewormanager/internal/repository/postgres/subjectnote"
	userRepo "homewormanager/internal/repository/postgres/user"
	userRedisRepo "homewormanager/internal/repository/redis/user"
	adminService "homewormanager/internal/service/admin"
	classService "homewormanager/internal/service/class"
	groupService "homewormanager/internal/service/group"
	homeworkService "homewormanager/internal/service/homework"
	homeworkFilesService "homewormanager/internal/service/homework_files"
	homeworkStatusService "homewormanager/internal/service/homework_status"
	moderatorService "homewormanager/internal/service/moderator"
	scheduleService "homewormanager/internal/service/schedule"
	subjectService "homewormanager/internal/service/subject"
	subjectNoteService "homewormanager/internal/service/subjectnote"
	userService "homewormanager/internal/service/user"
	"homewormanager/internal/tx_manager"
	postgres "homewormanager/pkg/db/postgres"
	"homewormanager/pkg/db/redis"
	"homewormanager/pkg/logger"
	"homewormanager/pkg/metrics"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func createApp() {

	//config
	cfg := config.NewCfg()

	//logger
	logger := logger.InitLogger(cfg)

	//metrics
	//TODO подумать про то нужно ли здесь что то возвращать
	metrics.StartMetricsServer(&cfg.MetricsConfig, logger)

	//database
	postgresDb := postgres.Connect(&cfg.PGConfig)
	redisClient := redis.Connect(&cfg.RedisConfig)

	//s3
	s3Client := s3.NewS3Client(&cfg.S3Config)

	//txmanager
	txmanager := tx_manager.NewTxManager(postgresDb)

	//Repo
	userRepo := userRepo.NewUserRepo(txmanager)
	userRedisRepo := userRedisRepo.NewUserRepo(redisClient, cfg)
	groupRepo := groupRepo.NewGroupRepo(txmanager)
	classRepo := classRepo.NewClassRepo(txmanager)
	subjectRepo := subjectRepo.NewSubjectRepo(txmanager)
	subjectNoteRepo := subjectNoteRepo.NewSubjectNoteRepo(txmanager)
	homeworkRepo := homeworkRepo.NewHomeworkRepo(txmanager)
	homeworkStatusRepo := homeworkStatusRepo.NewHomeworkStatusRepo(txmanager)
	homeworkFilesRepo := homeworkFilesRepo.NewHomeworkFilesRepo(txmanager)

	//Service
	userService := userService.NewUserService(userRepo, userRedisRepo, groupRepo, cfg)

	groupService := groupService.NewGroupService(groupRepo)

	subjectService := subjectService.NewSubjectService(subjectRepo)

	subjectNoteService := subjectNoteService.NewSubjectNoteService(subjectNoteRepo, subjectService)

	classService := classService.NewClassService(groupService, classRepo, subjectService, txmanager)

	homeworkStatusService := homeworkStatusService.NewHomeworkStatusService(homeworkStatusRepo)

	homeworkFilesService := homeworkFilesService.NewHomeworkFilesService(s3Client, homeworkFilesRepo, txmanager)

	homeworkService := homeworkService.NewHomeworkService(homeworkRepo, homeworkStatusService, txmanager)

	adminService := adminService.NewAdminService(groupService, classService, subjectService, homeworkService, userService, subjectNoteService, homeworkStatusService, txmanager)

	moderatorService := moderatorService.NewModeratorService(homeworkService, subjectNoteService, groupService, homeworkFilesService)

	scheduleService := scheduleService.NewScheduleService(classService, homeworkService, homeworkFilesService)

	//Handlers
	userHandler := userHandlers.NewUserHandler(cfg, userService)

	adminHandler := adminHandlers.NewAdminHandler(adminService)

	groupHandler := groupHandlers.NewGroupHandler(groupService)

	moderatorHandler := moderatorHandlers.NewModeratorHandler(moderatorService)

	scheduleHandler := scheduleHandlers.NewScheduleHandler(scheduleService)

	homeworkStatusHandler := homeworkStatusHandlers.NewHomeworkStatusHandler(homeworkStatusService)

	subjectHandler := subjectHandlers.NewSubjectHandler(subjectService)

	subjectNoteHandler := subjectNoteHandlers.NewSubjectNoteHandler(subjectNoteService)

	//fiber
	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: err_handler.ErrorHandler,
		BodyLimit:    20 * 1024 * 1024,
	})

	fiberApp.Use(cors.New(cors.Config{
		AllowOrigins:     fmt.Sprintf("https://%s:%s", cfg.FrontendHost, cfg.FrontendPort),
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS, PATCH",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))

	//middleware
	mw := middleware.NewMwManager(userRedisRepo, userRepo)

	fiberApp.Use(mw.Metrics(), mw.ErrorLogger(logger), mw.RequestLogger(logger))

	//groups
	userGroup := fiberApp.Group("/user")
	userHandlers.MapUserRoutes(userGroup, userHandler, mw)

	adminGroup := fiberApp.Group("/admin")
	adminHandlers.MapAdminRoutes(adminGroup, adminHandler, mw)

	groupGroup := fiberApp.Group("/group")
	groupHandlers.MapGroupRoutes(groupGroup, groupHandler)

	moderatorGroup := fiberApp.Group("/moderator")
	moderatorHandlers.MapModeratorRoutes(moderatorGroup, moderatorHandler, mw)

	scheduleGroup := fiberApp.Group("/schedule")
	scheduleHandlers.MapScheduleRoutes(scheduleGroup, scheduleHandler, mw)

	homeworkGroup := fiberApp.Group("/homework")
	homeworkStatusHandlers.MapHomeworkStatusRoutes(homeworkGroup, homeworkStatusHandler, mw)

	subjectGroup := fiberApp.Group("/subject")
	subjectHandlers.MapSubjectRoutes(subjectGroup, subjectHandler, mw)

	subjectNoteGroup := fiberApp.Group("/note")
	subjectNoteHandlers.MapSubjectNoteRoutes(subjectNoteGroup, subjectNoteHandler, mw)

	fiberApp.Get("/health", func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if err := postgresDb.PingContext(ctx); err != nil {
			return c.SendStatus(fiber.StatusServiceUnavailable)
		}
		return c.SendStatus(fiber.StatusOK)
	})
	fiberApp.Get("/livez", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusOK) })

	//fiber listen
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := fiberApp.Listen(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)); err != nil {
			panic(err)
		}
	}()

	<-exit
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := fiberApp.ShutdownWithContext(shutdownCtx); err != nil {
		panic(err)
	}

}
