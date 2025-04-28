package services

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/miloszbo/meals-finder/internal/models"
	repository "github.com/miloszbo/meals-finder/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

var key string = os.Getenv("APP_JWT_KEY")

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

func (s *UserService) GenerateJWT(username string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": username,
			"exp": time.Now().Add(24 * time.Hour).Unix(),
			"iat": time.Now().Unix(),
		})
	return t.SignedString(key)
}
