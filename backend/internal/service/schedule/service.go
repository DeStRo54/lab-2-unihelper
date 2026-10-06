package schedule

type Service struct {
	classService         ClassService
	homeworkService      HomeworkService
	homeworkFilesService HomeworkFilesService
}

func NewScheduleService(classService ClassService, homeworkService HomeworkService, homeworkFilesService HomeworkFilesService) *Service {
	return &Service{
		classService:         classService,
		homeworkService:      homeworkService,
		homeworkFilesService: homeworkFilesService,
	}
}
