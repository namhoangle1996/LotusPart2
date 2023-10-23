package model

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
