package data

import (
	"Greenlight/internal/validator"
	"database/sql"
	"errors"
	"time"
)

/* a. This model will have next fields: ID, CreatedAt, UpdatedAt, Name, Surname, Email,
PasswordHash, Role, Activated, Version.

*/

type User struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserName    string    `json:"user_name"`
	UserSurname string    `json:"user_surname"`
	Email       string    `json:"email"`

	PasswordHash string   `json:"password_hash"`
	Role         []string `json:"role"`
	Activated    bool     `json:"activated"`

	Version int32 `json:"version"`
}

type UsersModel struct {
	DB *sql.DB
}

func ValidateUser(v *validator.Validator, module *User) {
	v.Check(module.UserName != "", "user_name", "must be provided")
	v.Check(len(module.UserName) <= 500, "user_name", "must not be more than 500 bytes long")

	v.Check(module.UserSurname != "", "user_surname", "must be provided")
	v.Check(len(module.UserSurname) <= 500, "user_surname", "must not be more than 500 bytes long")

	v.Check(module.Email != "", "email", "must be provided")
	v.Check(len(module.Email) <= 500, "email", "must not be more than 500 bytes long")

	v.Check(len(module.Role) >= 1, "role", "must contain at least 1 role")
}

// method for inserting info to table
func (m UsersModel) Insert(module *User) error {
	//SQL query
	query := `
INSERT INTO user_info (user_name, user_surname, email, role)
VALUES ($1, $2, $3)
RETURNING id, created_at, updated_at, version`

	args := []any{module.UserName, module.UserSurname, module.Email, module.Role}

	return m.DB.QueryRow(query, args...).Scan(&module.ID, &module.CreatedAt, &module.UpdatedAt, &module.Version)
}

// method for fetching
func (m UsersModel) Get(id int64) (*User, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	//SQL query
	query := `SELECT id, created_at, updated_at, user_name,  user_surname, email, role, version 
FROM module_info 
WHERE id =$1`

	var module User

	err := m.DB.QueryRow(query, id).Scan(
		&module.ID,
		&module.CreatedAt,
		&module.UpdatedAt,
		&module.UserName,
		&module.UserSurname,
		&module.Email,
		&module.Role,
		&module.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}

	}

	return &module, nil

}

// method for updating
func (m UsersModel) Update(module *User) error {
	query := `
  UPDATE user_info
  SET  user_name =$1, user_surname=$2, email=$3, role=$4, version = version + 1
  WHERE id=$4
  RETURNING version
`

	args := []any{
		module.UserName,
		module.UserSurname,
		module.Email,
		module.Role,
		module.ID,
	}

	return m.DB.QueryRow(query, args...).Scan(&module.Version)

}

// method for deleting
func (m UsersModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
  DELETE FROM user_info
  WHERE id = $1
`

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	RowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if RowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
