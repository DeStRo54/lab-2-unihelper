package moderator

type Service struct {
	homeworkService     HomeworkService
	subjectNoteService  SubjectNoteService
	groupService        GroupService
	homeworkFileService HomeworkFileService
}

func NewModeratorService(homeworkService HomeworkService, subjectNoteService SubjectNoteService, groupService GroupService, homeworkFileService HomeworkFileService) *Service {
	return &Service{
		homeworkService:     homeworkService,
		subjectNoteService:  subjectNoteService,
		groupService:        groupService,
		homeworkFileService: homeworkFileService,
	}
}
