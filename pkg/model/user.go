package model

type User struct {
	ID       int    `json:"id" gorm:"primaryKey; autoIncrement" `
	UserName string `json:"user_name" gorm:"not null; unique; index"`
	Password string `json:"password"  gorm:"not null"`
}
