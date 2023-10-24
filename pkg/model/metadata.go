package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgtype"
	"time"
)

type User struct {
	ID       uint64 `json:"id" gorm:"primaryKey; autoIncrement" `
	UserName string `json:"user_name" gorm:"not null; unique; index"`
	Password string `json:"password"  gorm:"not null"`
}

func (t *User) TableName() string {
	return "user"
}

type Auth struct {
	ID     uint64 `json:"id" gorm:"primary_key;auto_increment" `
	UserID uint64 `json:"user_id" gorm:"not null" `
}

func (t *Auth) TableName() string {
	return "auth"
}

type AccessTokenClaim struct {
	jwt.StandardClaims
	AuthID uint64 `json:"authId" `
	UserID uint64 `json:"userId" `
}

type File struct {
	ID          uint64       `json:"id" gorm:"primaryKey; autoIncrement" `
	UserId      uint64       `json:"user_id" `
	Size        uint64       `json:"size" `
	ContentType string       `json:"content_type" `
	HttpInfo    pgtype.JSONB `json:"http_info" `
	CreatedAt   time.Time    `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (t *File) TableName() string {
	return "file"
}
