package repo

import (
	"LotusPart2/pkg/model"
	"context"
	"gorm.io/gorm"
)

func NewPGRepo(db *gorm.DB) PGInterface {
	return &RepoPG{db: db}
}

type PGInterface interface {
	CreateUser(ctx context.Context, req *model.User) error
	GetUserByUserName(ctx context.Context, userName string) (*model.User, error)
	GetUserByUserId(ctx context.Context, userId int) (*model.User, error)
}

type RepoPG struct {
	db    *gorm.DB
	debug bool
}

func (r *RepoPG) DB() *gorm.DB {
	return r.db
}
