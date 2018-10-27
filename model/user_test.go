package model

import (
	"reflect"
	"testing"

	"github.com/go-xorm/xorm"
	"github.com/sysu-go-online/public-service/types"
)

func TestUser_TableName(t *testing.T) {
	tests := []struct {
		name string
		u    User
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.TableName(); got != tt.want {
				t.Errorf("User.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Insert(t *testing.T) {
	type args struct {
		session *xorm.Session
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.Insert(tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_AddUserHome(t *testing.T) {
	tests := []struct {
		name    string
		u       *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.AddUserHome(); (err != nil) != tt.wantErr {
				t.Errorf("User.AddUserHome() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_GetWithEmail(t *testing.T) {
	type args struct {
		session *xorm.Session
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.GetWithEmail(tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.GetWithEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User.GetWithEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetWithUsername(t *testing.T) {
	type args struct {
		session *xorm.Session
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.GetWithUsername(tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.GetWithUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User.GetWithUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetWithUserID(t *testing.T) {
	type args struct {
		session *xorm.Session
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.GetWithUserID(tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.GetWithUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User.GetWithUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFileStructure(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *types.FileStructure
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFileStructure(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileStructure() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFileStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}
