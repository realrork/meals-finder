package services

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/miloszbo/meals-finder/internal/models"
	repository "github.com/miloszbo/meals-finder/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	DbConn *pgx.Conn
	Repo   *repository.Queries
}

func NewUserService(conn *pgx.Conn) UserService {
	return UserService{
		DbConn: conn,
		Repo:   repository.New(conn),
	}
}

func (s *UserService) LoginUser(ctx context.Context, loginData *models.LoginUserRequest) error {
	user, err := s.Repo.LoginUserWithUsername(ctx, loginData.Login)
	if err != nil {
		return errors.New("Wrong username or password.")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Passwdhash), []byte(loginData.Password)); err != nil {
		return errors.New("Wrong username or password")
	}

	return nil
}
