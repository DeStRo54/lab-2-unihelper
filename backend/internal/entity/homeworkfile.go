package entity

import "time"

type HomeworkFile struct {
	FileID     FileID
	HomeworkID HomeworkID
	GroupID    GroupID
	FileName   string
	FileURL    string
	Key        string
	CreatedAt  time.Time
}
