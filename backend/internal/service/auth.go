package service

import (
	"context"
	"errors"
	"knowledge/internal/auth"
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

func Login(ctx context.Context,username,password string)(token string,err error){
	//查询用户是否存在
	user, err := model.SelectByusername(ctx, username)
	if err!=nil {
		return "",errors.New("登录用户不存在")
	}
	//校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("密码不正确，请检查密码")
	}
	return auth.GenerateJwtToken(username,user.ID)

}
