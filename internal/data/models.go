package data

import (
	"database/sql"
	"errors"
)

// custom ErrRecordNotFound error.
var (
	ErrRecordNotFound = errors.New("record not found")
)

// Models struct which wraps the ModuleModel.
type Models struct {
	Modules     ModuleModel
	TeacherInfo TeachersModel
	UserInfo    UsersModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Modules:     ModuleModel{DB: db},
		TeacherInfo: TeachersModel{DB: db},
		UserInfo:    UsersModel{DB: db},
	}
}
