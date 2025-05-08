package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/miloszbo/meals-finder/internal/handlers"
	repository "github.com/miloszbo/meals-finder/internal/repositories"
	"github.com/miloszbo/meals-finder/internal/server"
	"github.com/miloszbo/meals-finder/internal/services"
)

type LoginTestStruct struct {
	Name  string
	Input string
	Want  int
}

func TestLoginUserMock(t *testing.T) {
	var tests []LoginTestStruct = []LoginTestStruct{
		{"No json pass", "", http.StatusBadRequest},
		{"Pass only login", `{"login":"david"}`, http.StatusBadRequest},
		{"Pass only password", `{"password":"D8Dsadsvd"}`, http.StatusBadRequest},
		{"Pass login and password", `{"login":"tomas", "password":"DSA43fFDD"}`, http.StatusOK},
		{"Pass empty login and empty password", `{"login":"","password":""}`, http.StatusBadRequest},
		{"Pass empty login and password", `{"login":"", "password":"F4DSA654gf"}`, http.StatusBadRequest},
		{"Pass login and empty password", `{"login":"eminem","password":""}`, http.StatusBadRequest},
	}

	handler := handlers.UserHandler{
		UserService: &services.MockUserService{},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/user/login", bytes.NewBufferString(tt.Input))
			resp := httptest.NewRecorder()
			handler.LoginUser(resp, req)
			if resp.Code != tt.Want {
				t.Errorf("got %v, want %v", resp.Code, tt.Want)
			} else {
				res := resp.Result()
				defer res.Body.Close()
				cookies := res.Cookies()
				if len(cookies) != 0 {
					fmt.Println(tt.Name)
					for _, c := range cookies {
						fmt.Printf("Cookie: %s\nValue: %s\n", c.Name, c.Value)
					}
				}
			}
		})
	}
}

// Requires first register user test
func TestLoginUserIntegration(t *testing.T) {
	var tests []LoginTestStruct = []LoginTestStruct{
		{"No json pass", "", http.StatusBadRequest},
		{"Pass only login", `{"login":"david"}`, http.StatusBadRequest},
		{"Pass only password", `{"password":"D8Dsadsvd"}`, http.StatusBadRequest},
		{"Pass login and password", `{"login":"tomas", "password":"DSA43fFDD"}`, http.StatusOK},
		{"Pass empty login and empty password", `{"login":"","password":""}`, http.StatusBadRequest},
		{"Pass empty login and password", `{"login":"", "password":"F4DSA654gf"}`, http.StatusBadRequest},
		{"Pass login and empty password", `{"login":"eminem","password":""}`, http.StatusBadRequest},
	}

	handler := handlers.UserHandler{
		UserService: &services.BaseUserService{
			DbConn: server.NewConnectionTest(),
			Repo:   repository.New(server.NewConnectionTest()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/user/login", bytes.NewBufferString(tt.Input))
			resp := httptest.NewRecorder()
			handler.LoginUser(resp, req)
			if resp.Code != tt.Want {
				t.Errorf("got %v, want %v", resp.Code, tt.Want)
			} else {
				res := resp.Result()
				defer res.Body.Close()
				cookies := res.Cookies()
				if len(cookies) != 0 {
					fmt.Println(tt.Name)
					for _, c := range cookies {
						fmt.Printf("Cookie: %s\nValue: %s\n", c.Name, c.Value)
					}
				}
			}
		})
	}
}
