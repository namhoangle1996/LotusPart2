package service

import (
	"LotusPart2/pkg/constant"
	"LotusPart2/pkg/model"
	"LotusPart2/pkg/repo"
	"context"
	"gitlab.com/goxp/cloud0/ginext"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserService struct {
	repo repo.PGInterface
}

func (s *UserService) Register(ctx context.Context, req model.RegisterRequest) (res *model.User, err error) {
	_, err = s.repo.GetUserByUserName(ctx, req.UserName)
	if err == nil {
		return nil, ginext.NewError(http.StatusBadRequest, constant.ERR_DUPLICATE_ACC_USERNAME)
	}

	bytesPw, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		UserName: req.UserName,
		Password: string(bytesPw),
	}

	err = s.repo.CreateUser(ctx, newUser)
	return newUser, err
}

func NewUserService(repo repo.PGInterface) UserInterface {
	return &UserService{repo: repo}
}

type UserInterface interface {
	Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error)
	Register(ctx context.Context, req model.RegisterRequest) (*model.User, error)
}

func (s *UserService) Login(ctx context.Context, req model.LoginRequest) (res *model.LoginResponse, err error) {
	getUser, err := s.repo.GetUserByUserName(ctx, req.UserName)
	if err == nil {
		return nil, ginext.NewError(http.StatusBadRequest, constant.ERR_NOT_EXIST_ACC_USERNAME)
	}

	if !checkPasswordHash(req.Password, getUser.Password) {
		return nil, ginext.NewError(http.StatusForbidden, constant.ERR_INCORRECT_ACC_PASS)
	}

	return res, err

}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//func createToken(authD AuthDetails) (string, error) {
//	claims := jwt.MapClaims{}
//	claims["authorized"] = true
//	claims["auth_uuid"] = authD.AuthUuid
//	claims["user_id"] = authD.UserId
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	return token.SignedString([]byte(os.Getenv("API_SECRET")))
//}
