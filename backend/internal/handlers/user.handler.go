package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/miloszbo/meals-finder/internal/models"
	"github.com/miloszbo/meals-finder/internal/services"
)

type UserHandler struct {
	UserService services.UserService
}

func NewUserHandler(conn *pgx.Conn) UserHandler {
	return UserHandler{
		UserService: services.NewUserService(conn),
	}
}

func (u *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	loginData := &models.LoginUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(loginData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := u.UserService.LoginUser(ctx, loginData); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
