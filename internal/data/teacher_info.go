package data

import "database/sql"

type TeacherInfo struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	ModuleID int    `json:"module_id"`
}

type TeachersModel struct {
	DB *sql.DB
}

func (m TeachersModel) GetAll() ([]*TeacherInfo, error) {
	query := `SELECT email, module_id FROM teachers_info`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teachers []*TeacherInfo

	for rows.Next() {
		teacher := &TeacherInfo{}
		if err := rows.Scan(
			&teacher.Email,
			&teacher.ModuleID,
		); err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return teachers, nil
}
