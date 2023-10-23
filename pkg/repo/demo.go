package repo

import (
	"LotusPart2/pkg/model"
	"context"
	"gitlab.com/goxp/cloud0/ginext"
	"gitlab.com/goxp/cloud0/logger"
	"net/http"
)

func (r *RepoPG) CreateUser(ctx context.Context, req *model.User) error {
	log := logger.WithCtx(ctx, "RepoPG.CreateUser")

	if err := r.db.WithContext(ctx).Create(req).Error; err != nil {
		log.WithError(err).Error("Error when call func CreateUser")
		return ginext.NewError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (r *RepoPG) GetUserByUserName(ctx context.Context, userName string) (rs *model.User, err error) {
	log := logger.WithCtx(ctx, "RepoPG.GetUserByUserName")

	if err := r.db.Where("user_name =? ", userName).Take(&rs).Error; err != nil {
		log.WithError(err).Error("Error when call func GetUserByUserName")
		return rs, err
	}

	return rs, nil
}

func (r *RepoPG) GetUserByUserId(ctx context.Context, userId int) (rs *model.User, err error) {
	log := logger.WithCtx(ctx, "RepoPG.GetUserByUserId")

	if err := r.db.Where("id =? ", userId).Take(&rs).Error; err != nil {
		log.WithError(err).Error("Error when call func GetUserByUserId")
		return rs, err
	}

	return rs, nil
}
