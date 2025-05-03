package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/miloszbo/meals-finder/internal/handlers"
	"github.com/miloszbo/meals-finder/internal/services"
)

func TestLoginUser(t *testing.T) {
	body := `{"login":"amanda", "password":"DSAEWQ123"}`
	req := httptest.NewRequest(http.MethodGet, "/user/login", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler := handlers.UserHandler{
		UserService: &services.MockUserService{},
	}

	handler.LoginUser(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %v", res.StatusCode)
	}

	cookies := res.Cookies()
	if len(cookies) == 0 {
		t.Fatal("Expected cookie to be set, got none")
	}
	for _, c := range cookies {
		fmt.Println(c.String())
	}
}
