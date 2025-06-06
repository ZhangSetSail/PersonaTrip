package models

import (
	"time"
)

// UserMySQL 用户模型（MySQL）
type UserMySQL struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"user_id" gorm:"size:100;not null;unique"` // 新增 UserID 字段
	Username  string    `json:"username" gorm:"size:50;not null;unique"`
	Email     string    `json:"email" gorm:"size:100;not null;unique"`
	Password  string    `json:"-" gorm:"size:100;not null"` // 不在JSON中返回密码
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *UserMySQL) TableName() string {
	return "users"
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string    `json:"token"`
	User  UserMySQL `json:"user"`
}

// TokenClaims JWT令牌声明
type TokenClaims struct {
	UserID uint `json:"user_id"`
}
