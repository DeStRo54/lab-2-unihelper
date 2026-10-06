package entity

const (
	RoleUser Role = iota + 1
	RoleGroupModerator
	RoleGlobalAdmin
)

const (
	SessionKey = "session_key"
	Claims     = "claims"
)

const (
	CategoryLection      = "ЛК"
	CategoryPractice     = "ПР"
	CategoryLab          = "ЛАБ"
	CategoryCredit       = "ЗАЧ"
	CategoryConsultation = "Конс"
	CategoryExam         = "Э"
)

var CategoryToNumber = map[string]ClassCategory{
	CategoryLection:      1,
	CategoryPractice:     2,
	CategoryLab:          3,
	CategoryCredit:       4,
	CategoryConsultation: 5,
	CategoryExam:         6,
}

var NumberToCategory = map[ClassCategory]string{
	1: CategoryLection,
	2: CategoryPractice,
	3: CategoryLab,
	4: CategoryCredit,
	5: CategoryConsultation,
	6: CategoryExam,
}

var PermittedExt = map[string]bool{
	".pdf":  true,
	".doc":  true,
	".docx": true,
	".png":  true,
	".jpg":  true,
	".jpeg": true,
	".rar":  true,
	".7z":   true,
	".zip":  true,
	".ppt":  true,
	".pptx": true,
	".odt":  true,
	".xlsx": true,
	".xls":  true,
	".txt":  true,
}
