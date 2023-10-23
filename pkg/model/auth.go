package model

type Auth struct {
	ID       uint64 `json:"id" gorm:"primary_key;auto_increment" `
	UserID   uint64 `json:"user_id" gorm:"not null" `
	AuthUUID string `json:"auth_uuid" gorm:"size:255;not null" `
}
