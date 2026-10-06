package main

import (
	"github.com/robfig/cron/v3"
	"homewormanager/internal/config"
	"homewormanager/internal/cron/schedule_updater"
	classRepo "homewormanager/internal/repository/postgres/class"
	groupRepo "homewormanager/internal/repository/postgres/group"
	homeworkRepo "homewormanager/internal/repository/postgres/homework"
	homeworkStatusRepo "homewormanager/internal/repository/postgres/homework_status"
	subjectRepo "homewormanager/internal/repository/postgres/subject"
	subjectNoteRepo "homewormanager/internal/repository/postgres/subjectnote"
	userRepo "homewormanager/internal/repository/postgres/user"
	userRedisRepo "homewormanager/internal/repository/redis/user"
	adminService "homewormanager/internal/service/admin"
	classService "homewormanager/internal/service/class"
	groupService "homewormanager/internal/service/group"
	homeworkService "homewormanager/internal/service/homework"
	homeworkStatusService "homewormanager/internal/service/homework_status"
	subjectService "homewormanager/internal/service/subject"
	subjectNoteService "homewormanager/internal/service/subjectnote"
	userService "homewormanager/internal/service/user"
	"homewormanager/internal/tx_manager"
	"homewormanager/pkg/db/postgres"
	"homewormanager/pkg/db/redis"
	"homewormanager/pkg/logger"
	"time"
)

func main() {
	//config
	cfg := config.NewCfg()

	//database
	postgresDb := postgres.Connect(&cfg.PGConfig)
	redisClient := redis.Connect(&cfg.RedisConfig)

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

	//logger
	logger := logger.InitLogger(cfg)

	//Service
	groupService := groupService.NewGroupService(groupRepo)

	userService := userService.NewUserService(userRepo, userRedisRepo, groupService, cfg)

	subjectService := subjectService.NewSubjectService(subjectRepo)

	subjectNoteService := subjectNoteService.NewSubjectNoteService(subjectNoteRepo, subjectService)

	classService := classService.NewClassService(groupService, classRepo, subjectService, txmanager)

	homeworkStatusService := homeworkStatusService.NewHomeworkStatusService(homeworkStatusRepo)

	homeworkService := homeworkService.NewHomeworkService(homeworkRepo, homeworkStatusService, txmanager)

	adminService := adminService.NewAdminService(groupService, classService, subjectService, homeworkService, userService, subjectNoteService, homeworkStatusService, txmanager)
	job := schedule_updater.NewCronJob(adminService, logger)

	c := cron.New(cron.WithLocation(time.Local))
	_, err := c.AddFunc("0 0 * * *", job.Run)
	logger.Info("Worker start")
	if err != nil {
		panic(err)
	}
	c.Start()
	defer c.Stop()

	select {}
}
