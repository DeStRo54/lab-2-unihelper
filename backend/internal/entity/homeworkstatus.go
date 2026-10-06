package entity

type HomeworkStatus struct {
	ID         StatusID
	UserID     UserID
	HomeworkID HomeworkID
	Status     bool
}
