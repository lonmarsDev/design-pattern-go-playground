package repository

import "github.com/realpamisa/RestAPI/internal/api/model"

type UserRepository interface {
	Save(user model.User) (model.User, error)
	Update(id string, user model.UpdateUser) (bool, error)
	FindByID(id string) (model.User, error)
	Delete(id string) (bool, error)
	FindAll() ([]model.User, error)
}
