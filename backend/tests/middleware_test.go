package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/miloszbo/meals-finder/internal/middlewares"
)

var key []byte = []byte(os.Getenv("APP_JWT_KEY"))

func createTestToken(expired bool) string {
	exp := time.Now().Add(time.Hour).Unix()

	if expired {
		exp = time.Now().Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "testToken",
		"exp": exp,
		"iat": time.Now().Unix(),
	})

	tokenString, _ := token.SignedString(key)
	return tokenString
}

type AuthTestStruct struct {
	Name  string
	Input string
	Want  int
}

func TestAuthentication(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	})

	var tests []AuthTestStruct = []AuthTestStruct{
		{"No token", "", 401},
		{"Expired token", "Bearer " + createTestToken(true), http.StatusUnauthorized},
		{"Valid token", "Bearer " + createTestToken(false), http.StatusOK},
		{"Wrong signature token", "Bearer " + createTestToken(false) + "R", http.StatusUnauthorized},
		{"Expired and wrong signature token", "Bearer " + createTestToken(true) + "R", http.StatusUnauthorized},
	}

	handlerTest := middlewares.Authentication(handler)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			resp := httptest.NewRecorder()
			req.Header.Set("Authorization", tt.Input)
			handlerTest.ServeHTTP(resp, req)
			if resp.Code != tt.Want {
				t.Errorf("got %v, want %v", resp.Code, tt.Want)
			}
		})
	}
}
