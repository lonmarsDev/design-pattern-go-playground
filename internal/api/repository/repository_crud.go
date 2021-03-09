package repository

import (
	"github.com/realpamisa/RestAPI/internal/api/model"
	"gorm.io/gorm"
)

type userRepositoryCrud struct {
	db *gorm.DB
}

func NewRepositoryUserCRUD(db *gorm.DB) *userRepositoryCrud {
	return &userRepositoryCrud{db}
}

func (u *userRepositoryCrud) Save(user model.User) (model.User, error) {
	if !u.db.Migrator().HasTable(&user) {
		u.db.Migrator().CreateTable(&user)
	}
	err := u.db.Debug().Model(&user).Create(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepositoryCrud) Update(id string, user model.UpdateUser) (bool, error) {
	var userTable model.User
	if !u.db.Migrator().HasTable(&user) {
		u.db.Migrator().CreateTable(&user)
	}
	err := u.db.Model(&userTable).Where("id = ?", id).Save(user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
