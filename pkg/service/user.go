package service

import (
	"LotusPart2/conf"
	"LotusPart2/pkg/constant"
	"LotusPart2/pkg/model"
	"LotusPart2/pkg/repo"
	"context"
	"github.com/dgrijalva/jwt-go"
	"gitlab.com/goxp/cloud0/ginext"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type UserService struct {
	repo repo.PGInterface
}

func (s *UserService) Logout(ctx context.Context, userId int64) error {
	return s.repo.DeleteAuthByUserId(ctx, userId)
}

func (s *UserService) UploadFile(ctx context.Context, userId, authId int64) error {
	_, err := s.repo.GetAuthByIdAndUserId(ctx, userId, authId)
	if err != nil {
		return ginext.NewError(http.StatusForbidden, "Token is invalid")
	}

	//s.repo.saveFile()
	return nil
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
	Logout(ctx context.Context, userId int64) error

	UploadFile(ctx context.Context, userIdHeader, authId int64) error
}

func (s *UserService) Login(ctx context.Context, req model.LoginRequest) (res *model.LoginResponse, err error) {
	getUser, err := s.repo.GetUserByUserName(ctx, req.UserName)
	if err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, constant.ERR_NOT_EXIST_ACC_USERNAME)
	}

	if !checkPasswordHash(req.Password, getUser.Password) {
		return nil, ginext.NewError(http.StatusForbidden, constant.ERR_INCORRECT_ACC_PASS)
	}

	authData := model.Auth{
		UserID: getUser.ID,
	}
	err = s.repo.CreateAuth(ctx, &authData)
	if err != nil {
		return
	}

	token, err := createToken(authData)
	if err != nil {
		return nil, err
	}

	var resp = model.LoginResponse{
		Token: token,
	}

	return &resp, err

}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func createToken(authD model.Auth) (string, error) {

	expiredAt := time.Now().Add(time.Minute * time.Duration(5)).Unix()
	claims := &jwt.StandardClaims{
		Audience:  "audience",
		ExpiresAt: expiredAt,
		Issuer:    "issuer",
		Subject:   "subject",
	}

	claim := model.AccessTokenClaim{
		StandardClaims: *claims,
		AuthID:         authD.ID,
		UserID:         authD.UserID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(conf.LoadEnv().API_SECRET))
}
