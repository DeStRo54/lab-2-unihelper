package homework_status

type Service struct {
	homeworkStatusRepo HomeworkStatusRepo
}

func NewHomeworkStatusService(homeworkStatusRepo HomeworkStatusRepo) *Service {
	return &Service{
		homeworkStatusRepo: homeworkStatusRepo,
	}
}
