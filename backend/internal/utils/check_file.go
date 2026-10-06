package utils

import (
	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
	"mime/multipart"
	"path/filepath"
	"strings"
)

func CheckFile(fileHeader *multipart.FileHeader) error {
	if fileHeader.Size > 20*1024*1024 {
		return errs.FileTooLarge
	}
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !entity.PermittedExt[ext] {
		return errs.InvalidFileType
	}
	return nil
}
