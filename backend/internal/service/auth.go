package service

import (
	"context"
	"knowledge/internal/model"

	"golang.org/x/crypto/bcrypt"
)

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func Register(txt context.Context, req *RegisterReq)(error) {

	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := model.User{
		Username:     req.Username,
		PasswordHash: string(PasswordHash),
		Email:        req.Email,
		Status: "active",
	}
	if err:=model.CreateUser(txt,&user);err!=nil {
		return err
	} 
	return nil
}
