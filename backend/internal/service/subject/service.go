package subject

type Service struct {
	subjectRepo SubjectRepo
}

func NewSubjectService(subjectRepo SubjectRepo) *Service {
	return &Service{
		subjectRepo: subjectRepo,
	}
}
