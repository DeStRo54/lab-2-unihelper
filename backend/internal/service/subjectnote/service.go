package subjectnote

type Service struct {
	subjectNoteRepo SubjectNoteRepo
	subjectService  SubjectService
}

func NewSubjectNoteService(subjectNoteRepo SubjectNoteRepo, subjectService SubjectService) *Service {
	return &Service{
		subjectNoteRepo: subjectNoteRepo,
		subjectService:  subjectService,
	}
}
