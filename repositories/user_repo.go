package repositories

import models "demo-api/model"

// cac methods lam viec voi db
type UserRepo interface {
	Select() ([]models.User, error) // pointer -> tranh tao ra version copy User model
	Insert(u models.User) error
}
