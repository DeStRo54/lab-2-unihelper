package errs

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
)

var (
	UserNotFound           = errors.New("user doesn't exist")
	GroupNotFound          = errors.New("group doesn't exist")
	SubjectsNotFound       = errors.New("subjects doesn't exist")
	ClassesNotFound        = errors.New("classes doesn't exist")
	HomeworksNotFound      = errors.New("homeworks doesn't exist")
	HomeworkStatusNotFound = errors.New("homeworkStatus doesn't exist")
	SubjectNotesNotFound   = errors.New("subjectNote doesn't exist")
	FileNotFound           = errors.New("file doesn't exist")
)

var (
	UserExists         = errors.New("user already exists")
	GroupExists        = errors.New("group already exists")
	ErrInvalidPassword = errors.New("invalid password provided")
)
var (
	InvalidRegisterKey = errors.New("invalid register key provided")
)
var (
	NoIcal = errors.New("no calendar")
)
var (
	NoUserRole = errors.New("no user role")
)

var (
	FileTooLarge    = errors.New("file too large")
	InvalidFileType = errors.New("invalid file type")
	TooManyFiles    = errors.New("too many files")
)

var (
	notFoundErrors = []error{
		sql.ErrNoRows,
		UserNotFound,
		GroupNotFound,
		ClassesNotFound,
		SubjectNotesNotFound,
		SubjectsNotFound,
		HomeworksNotFound,
		HomeworkStatusNotFound,
	}
	badRequestErrors = []error{
		UserExists,
		GroupExists,
		ErrInvalidPassword,
		InvalidRegisterKey,
		FileTooLarge,
	}
)

func isNotFoundError(err error) bool {
	for _, e := range notFoundErrors {
		if errors.Is(e, err) {
			return true
		}
	}
	return false
}

func isBadRequest(err error) bool {
	for _, e := range badRequestErrors {
		if errors.Is(e, err) {
			return true
		}
	}
	return false
}

func GetStatusFromError(err error) int {
	var ferr *fiber.Error
	if errors.As(err, &ferr) {
		return ferr.Code
	}

	if isNotFoundError(err) {
		return fiber.StatusNotFound
	}
	if isBadRequest(err) {
		return fiber.StatusBadRequest
	}

	//any unspecified error

	return fiber.StatusInternalServerError
}
