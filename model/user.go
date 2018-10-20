package model

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/sysu-go-online/public-service/tools"
	"github.com/sysu-go-online/public-service/types"
)

// User correspond user table in mysql
type User struct {
	ID         int        `xorm:"pk autoincr 'id'"`
	Username   string     `xorm:"notnull unique"`
	Email      string     `xorm:"notnull unique"`
	Password   string     `xorm:"notnull"`
	CreateTime *time.Time `xorm:"created"`
}

// TableName defines table name
func (u User) TableName() string {
	return "user"
}

// Insert a user to the table
func (u *User) Insert(session *xorm.Session) (int, error) {
	affected, err := session.InsertOne(u)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return int(affected), nil
}

// AddUserHome create user home directory
func (u *User) AddUserHome() error {
	userHome := path.Join("/home", u.Username)
	return os.MkdirAll(userHome, os.ModeDir)
}

// GetWithEmail get user with given email
func (u *User) GetWithEmail(session *xorm.Session) (bool, error) {
	email := u.Email
	return session.Where("email = ?", email).Get(u)
}

// GetWithUsername get user with given username
func (u *User) GetWithUsername(session *xorm.Session) (bool, error) {
	username := u.Username
	return session.Where("username = ?", username).Get(u)
}

// GetWithUserID get user with given id
func (u *User) GetWithUserID(session *xorm.Session) (bool, error) {
	id := u.ID
	return session.Where("id = ?", id).Get(u)
}

// GetFileStructure read file structure and return it
func GetFileStructure(username string) (*types.FileStructure, error) {
	// Get absolute path
	var err error
	absPath := filepath.Join("/", username)

	// Recurisively get file structure
	s, err := tools.Dfs(absPath, 0)
	if err != nil {
		return nil, err
	}
	// Add root content
	root := types.FileStructure{
		Name:       "",
		Type:       "dir",
		Children:   s,
		Root:       true,
		IsSelected: true,
	}
	return &root, nil
}
