package schedule

import (
	"time"

	"homewormanager/internal/entity"
)

type homework struct {
	HomeworkID   entity.HomeworkID `json:"homeworkID"`
	SubjectName  string            `json:"subjectName"`
	HomeworkText string            `json:"homeworkText"`
	IsCompleted  bool              `json:"isCompleted"`
	Files        []file            `json:"files"`
	DueDate      time.Time         `json:"dueDate"`
}

type file struct {
	FileID    entity.FileID
	FileName  string
	FileURL   string
	CreatedAt time.Time
}
type class struct {
	SubjectID      entity.SubjectID `json:"subjectId"`
	StartTime      time.Time        `json:"startTime"`
	EndTime        time.Time        `json:"endTime"`
	Summary        string           `json:"summary"`
	Description    string           `json:"description"`
	SemClassNumber int64            `json:"semClassNumber"`
	Category       string           `json:"category"`
	Location       string           `json:"location"`
}

type outputClass struct {
	Class    class      `json:"class"`
	Homework []homework `json:"homework"`
}

type scheduleDay struct {
	OutputClass         []outputClass `json:"outputClasses"`
	IndependentHomework []homework    `json:"independentHomeworks"`
}

func toClass(c entity.Class) class {
	return class{
		SubjectID:      c.SubjectID,
		StartTime:      c.StartTime,
		EndTime:        c.EndTime,
		Summary:        c.Summary,
		Description:    c.Description,
		SemClassNumber: c.SemClassNumber,
		Category:       entity.NumberToCategory[c.Category],
		Location:       c.Location,
	}
}
func toHomework(c []entity.Homework) []homework {
	result := make([]homework, len(c))
	for i, h := range c {
		result[i] = homework{
			HomeworkID:   h.HomeworkID,
			SubjectName:  h.SubjectName,
			HomeworkText: h.HomeworkText,
			IsCompleted:  h.IsCompleted,
			Files:        toFiles(h.Files),
			DueDate:      h.DueDate,
		}
	}
	return result
}

func toOutputClass(c []entity.OutputClass) []outputClass {
	result := make([]outputClass, len(c))
	for i, h := range c {
		result[i] = outputClass{
			Class:    toClass(h.Class),
			Homework: toHomework(h.Homework),
		}
	}
	return result
}

type homeworkDay struct {
	Homework []homework `json:"homework"`
}

func toFile(f entity.HomeworkFile) file {
	return file{
		FileID:    f.FileID,
		FileName:  f.FileName,
		FileURL:   f.FileURL,
		CreatedAt: f.CreatedAt,
	}
}

func toFiles(f []entity.HomeworkFile) []file {
	result := make([]file, len(f))
	for i := range f {
		result[i] = toFile(f[i])
	}
	return result
}
