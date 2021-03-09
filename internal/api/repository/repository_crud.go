package repository

import (
	"fmt"

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
	var userTable model.UpdateUser
	if !u.db.Migrator().HasTable(&user) {
		u.db.Migrator().CreateTable(&user)
	}
	if user.Firstname != nil {
		userTable.Firstname = user.Firstname
	}
	if user.Lastname != nil {
		userTable.Lastname = user.Lastname
	}
	if user.Email != nil {
		userTable.Email = user.Email
	}
	if user.Password != nil {
		userTable.Password = user.Password
	}

	err := u.db.Table("users").Where("id = ?", id).Updates(&userTable).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *userRepositoryCrud) FindByID(id string) (model.User, error) {

	var user model.User

	err := u.db.Model(&user).Where("id = ?", id).First(user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil

}

func (u *userRepositoryCrud) Delete(id string) (bool, error) {
	var user model.User

	err := u.db.Delete(user, "id= ", id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *userRepositoryCrud) FindAll() ([]model.User, error) {
	fmt.Println("check findall")
	var users []model.User
	u.db.Find(&users)
	return users, nil
}
