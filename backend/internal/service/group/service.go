package group

type Service struct {
	groupRepo GroupRepo
}

func NewGroupService(groupRepo GroupRepo) *Service {
	return &Service{
		groupRepo: groupRepo,
	}
}
