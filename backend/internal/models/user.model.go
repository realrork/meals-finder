package models

import "errors"

type LoginUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (lur *LoginUserRequest) Validate() error {
	if lur.Login == "" || lur.Password == "" {
		return errors.New("bad request")
	}
	return nil
}
