package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sysu-go-online/public-service/tools"
	"github.com/sysu-go-online/user-service/model"
)

// UserController is controller for user
type UserController struct {
	model.User
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserMessage stores user message
type UserMessage struct {
	Email string `json:"email"`
}

// CreateUserHandler handles user sign up
func CreateUserHandler(w http.ResponseWriter, r *http.Request) error {
	r.ParseForm()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	user := UserController{}
	if err := json.Unmarshal(body, &user); err != nil {
		return err
	}
	if ok := tools.CheckEmail(user.Email); !ok {
		w.WriteHeader(400)
		return nil
	}
	if ok := tools.CheckUsername(user.Username); !ok {
		w.WriteHeader(400)
		return nil
	}
	pass, err := tools.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.User.Email = user.Email
	user.User.Password = pass
	user.User.Username = user.Username
	user.Username = tools.GenerateUserName()

	// create session to add a user
	session := MysqlEngine.NewSession()
	affected, err := user.User.Insert(session)
	if err != nil {
		session.Rollback()
		return err
	}
	err = user.User.AddUserHome()
	if err != nil {
		session.Rollback()
		return err
	}
	session.Commit()
	if affected == 0 {
		w.WriteHeader(400)
		return nil
	}
	return nil
}

// GetUserMessageHandler handle user message query
func GetUserMessageHandler(w http.ResponseWriter, r *http.Request) error {
	username := mux.Vars(r)["username"]
	if ok := tools.CheckUsername(username); !ok {
		w.WriteHeader(400)
		return nil
	}
	u := UserController{}
	u.User.Username = username

	// create session to query a user
	session := MysqlEngine.NewSession()
	has, err := u.User.GetWithUsername(session)
	if err != nil {
		session.Rollback()
		return err
	}
	session.Commit()
	if !has {
		w.WriteHeader(204)
		return nil
	}
	um := UserMessage{u.User.Email}
	byteUM, err := json.Marshal(um)
	if err != nil {
		return err
	}
	w.Write(byteUM)
	return nil
}

// GetUserHomeStructureHandler returns home structure
func GetUserHomeStructureHandler(w http.ResponseWriter, r *http.Request) error {
	// get username from jwt
	ok, username := tools.GetUserNameFromToken(r.Header.Get("Authorization"), AuthRedisClient)
	if !ok {
		w.WriteHeader(401)
		return nil
	}

	// Get project information
	session := MysqlEngine.NewSession()
	u := model.User{Username: username}
	ok, err := u.GetWithUsername(session)
	if !ok {
		w.WriteHeader(400)
		return nil
	}
	if err != nil {
		return err
	}

	// Get file structure
	structure, err := model.GetFileStructure(username)
	if err != nil {
		return err
	}
	ret, err := json.Marshal(structure)
	if err != nil {
		return err
	}
	w.Write(ret)
	return nil
}
