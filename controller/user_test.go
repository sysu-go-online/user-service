package controller

import (
	"net/http"
	"testing"
)

func TestCreateUserHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUserHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("CreateUserHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserMessageHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetUserMessageHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("GetUserMessageHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserHomeStructureHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetUserHomeStructureHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("GetUserHomeStructureHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
