package controller

import (
	"knowledge/internal/apperror"
	"knowledge/internal/response"
	"knowledge/internal/service"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func Register(c *gin.Context) {

	var req service.RegisterReq
	err := c.ShouldBind(&req)
	if err != nil {
		_ = c.Error(apperror.New(http.StatusBadRequest, "请检查入参"))
		return
	}

	username := req.Username
	password := req.Password
	email := req.Email

	if username == "" {
		_ = c.Error(apperror.New(http.StatusBadRequest, "用户名不能为空"))
		return
	}
	if password == "" {
		_ = c.Error(apperror.New(http.StatusBadRequest, "密码不能为空"))
		return
	}
	if len(password) < 6 || len(password) > 18 {
		_ = c.Error(apperror.New(http.StatusBadRequest, "密码长度请控制在6~18之间"))
		return
	}
	if email == "" {
		_ = c.Error(apperror.New(http.StatusBadRequest, "邮箱不能为空"))
		return
	}
	if !IsValidEmail(email) {
		_ = c.Error(apperror.New(http.StatusBadRequest, "邮箱格式不正确"))
		return
	}

	if err := service.Register(c.Request.Context(), &req); err != nil {
		_ = c.Error(apperror.New(http.StatusInternalServerError, err.Error()))
		return
	}

	response.Success(c, "")

}
