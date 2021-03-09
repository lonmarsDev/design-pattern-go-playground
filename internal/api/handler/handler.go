package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/realpamisa/RestAPI/internal/api/model"
	"github.com/realpamisa/RestAPI/internal/api/repository"
	repo "github.com/realpamisa/RestAPI/internal/api/repository"

	"github.com/realpamisa/RestAPI/internal/api/response"
	"github.com/realpamisa/RestAPI/pkg/jwt"
	"github.com/realpamisa/RestAPI/pkg/mysql"
)

type Handler struct {
}

func Init() *Handler {
	return &Handler{}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	db, err := mysql.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
	}

	repo := repo.NewRepositoryUserCRUD(db)
	func(userRepository repository.UserRepository) {
		user, err := userRepository.FindByID(userID)
		if err != nil {
			response.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user))
		response.JSON(w, http.StatusCreated, user)
	}(repo)

}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := mysql.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
	}

	repo := repo.NewRepositoryUserCRUD(db)

	func(userRepository repository.UserRepository) {
		user, err := userRepository.FindAll()
		if err != nil {
			response.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user))
		response.JSON(w, http.StatusCreated, user)
	}(repo)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	db, err := mysql.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
	}
	hashPassword, err := jwt.PasswordHashAndSalt([]byte(user.Password))
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	newUser := model.User{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  *hashPassword,
	}
	repo := repo.NewRepositoryUserCRUD(db)
	func(userRepository repository.UserRepository) {
		user, err := userRepository.Save(newUser)
		if err != nil {
			response.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.Firstname))
		response.JSON(w, http.StatusCreated, user)
	}(repo)

}

//Update user from db
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user model.UpdateUser

	userID := chi.URLParam(r, "id")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	if user.Password != nil {
		hashPassword, err := jwt.PasswordHashAndSalt([]byte(*user.Password))
		if err != nil {
			response.ERROR(w, http.StatusUnprocessableEntity, err)
		}
		user = model.UpdateUser{
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
			Password:  hashPassword,
		}
	}

	db, err := mysql.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
	}

	repo := repo.NewRepositoryUserCRUD(db)
	func(userRepository repository.UserRepository) {
		user, err := userRepository.Update(userID, user)
		if err != nil {
			response.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user))
		response.JSON(w, http.StatusCreated, user)
	}(repo)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	db, err := mysql.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
	}

	repo := repo.NewRepositoryUserCRUD(db)
	func(userRepository repository.UserRepository) {
		user, err := userRepository.FindByID(userID)
		if err != nil {
			response.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user))
		response.JSON(w, http.StatusCreated, user)
	}(repo)
}

func Login(w http.ResponseWriter, r *http.Request) {

	var loginVar model.LoginVar

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(body, &loginVar)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	db, err := mysql.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
	}
	repo := repo.NewRepositoryUserCRUD(db)
	func(userRepository repository.UserRepository) {
		user, err := userRepository.Login(loginVar)
		if err != nil {
			response.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user))
		response.JSON(w, http.StatusCreated, user)
	}(repo)
}
