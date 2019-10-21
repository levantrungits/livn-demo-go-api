// Package repoimpl with Golang
// Designed by TRUNGLV
package repoimpl

import (
	"database/sql"
	models "demo-api/model"
	repo "demo-api/repositories"
	"fmt"
)

// Db Object UserRepoImpl
type UserRepoImpl struct {
	Db *sql.DB
}

// Contructopr NewUserRepo
func NewUserRepo(db *sql.DB) repo.UserRepo {
	return &UserRepoImpl{
		Db: db,
	}
}

// Public Select Function
func (u *UserRepoImpl) Select() ([]models.User, error) {
	users := make([]models.User, 0)

	rows, err := u.Db.Query("SELECT id, name, gender, email FROM users")
	if err != nil {
		return users, err
	}

	for rows.Next() {
		user := models.User{}
		// Select id, name, gender, email FROM public.users;
		err := rows.Scan(&user.ID, &user.Name, &user.Gender, &user.Email)
		if err != nil {
			break
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return users, err
	}

	return users, nil
}

// Public Insert Function
func (u *UserRepoImpl) Insert(user models.User) error {
	insertStatement := `
	INSERT INTO public.users(gender, id, name, email)
		VALUES ($1, $2, $3, $4)
	`
	_, err := u.Db.Exec(insertStatement, user.Gender, user.ID, user.Name, user.Email)
	if err != nil {
		return err
	}
	fmt.Println("Record added: ", user)
	return nil
}
