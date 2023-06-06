package model

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID         int64     `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Nick       string    `json:"nick,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
}

func (user *User) Run(etapa string) error {
	if err := user.valid(etapa); err != nil {
		return err
	}
	if err := user.Format(etapa); err != nil {
		return err
	}
	return nil
}

func (user *User) valid(etapa string) error {
	if user.Name == "" {
		return errors.New("username is required")
	}
	if user.Nick == "" {
		return errors.New("nick is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("email is not valid")
	}

	if etapa == "cadastro" && user.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (user *User) Format(etapa string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if etapa == "cadastro" {
		passHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(passHash)
	}
	return nil
}
