package handlers

import (
	"errors"
	"net/http"

	"github.com/miloszbo/meals-finder/internal/services"
)

var ErrBarRequest = errors.New("bad request")

func StatusFromError(err error) int {
	status := http.StatusInternalServerError

	switch err {
	case services.ErrUnauthorizedUser:
		status = http.StatusUnauthorized
	case services.ErrInternalFailure:
		status = http.StatusInternalServerError
	}

	return status
}
