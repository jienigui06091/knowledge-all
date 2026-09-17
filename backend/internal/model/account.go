package model

import (
	"context"
	"time"
)

const (
	UserStatusActive   = "active"
	UserStatusDisabled = "disabled"

	RoleCodeUser  = "user"
	RoleCodeAdmin = "admin"
)

// Role 对应 roles 表。
type Role struct {
	ID          int16   `json:"id"`
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Base

	UserRoles []UserRole `json:"-" gorm:"foreignKey:RoleID"`
}

func (Role) TableName() string {
	return "roles"
}

// User 对应 users 表。
type User struct {
	ID           int64      `json:"id"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"-"`
	DisplayName  string     `json:"display_name"`
	AvatarURL    *string    `json:"avatar_url"`
	Bio          *string    `json:"bio"`
	Status       string     `json:"status"`
	LastLoginAt  *time.Time `json:"last_login_at"`
	Base
	SoftDelete

	UserRoles []UserRole `json:"-" gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}

// UserRole 对应 user_roles 表，撤销角色时保留软删除历史。
type UserRole struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
	RoleID int16 `json:"role_id"`
	Base
	SoftDelete

	User User `json:"-" gorm:"foreignKey:UserID"`
	Role Role `json:"role" gorm:"foreignKey:RoleID"`
}

func (UserRole) TableName() string {
	return "user_roles"
}

func CreateUser(ctx context.Context, user *User) error {
	 return DB.WithContext(ctx).Create(user).Error
}
