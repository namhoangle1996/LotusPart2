package handler

import (
	"LotusPart2/pkg/model"
	"LotusPart2/pkg/service"
	"LotusPart2/pkg/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gitlab.com/goxp/cloud0/ginext"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type UserHandler struct {
	service  service.UserInterface
	validate *validator.Validate
}

func NewUserHandler(service service.UserInterface, validate *validator.Validate) *UserHandler {
	return &UserHandler{
		service:  service,
		validate: validate,
	}
}

// CreateMission
// @Tags User
// @Security ApiKeyAuth
// @Summary Login
// @Description Login
// @Accept  json
// @Produce  json
// @Param data body model.LoginRequest true "body data"
// @Success 200 {object} model.User
// @Router /api/v1/user/login [post]
func (h *UserHandler) Login(r *ginext.Request) (*ginext.Response, error) {
	var req model.LoginRequest
	if err := r.GinCtx.BindJSON(&req); err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}

	rs, err := h.service.Login(r.GinCtx, req)
	if err != nil {
		fmt.Println("err", err)
		r.GinCtx.JSON(http.StatusForbidden, err)
		return nil, nil
	}

	return ginext.NewResponseData(http.StatusOK, rs), nil
}

// Register
// @Tags User
// @Security ApiKeyAuth
// @Summary Register
// @Description Register
// @Accept  json
// @Produce  json
// @Param data body model.RegisterRequest true "body data"
// @Success 200 {object} model.User
// @Router /api/v1/user/register [post]
func (h *UserHandler) Register(r *ginext.Request) (*ginext.Response, error) {
	var req model.RegisterRequest
	if err := r.GinCtx.BindJSON(&req); err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}

	rs, err := h.service.Register(r.GinCtx, req)
	if err != nil {
		r.GinCtx.JSON(http.StatusForbidden, err)
		return nil, nil
	}

	return ginext.NewResponseData(http.StatusOK, rs), nil
}

// Upload file
// @Tags User
// @Security ApiKeyAuth
// @Summary // Upload file
// @Description // Upload file
// @Accept  json
// @Produce  json
// @Param data body model.RegisterRequest true "body data"
// @Success 200 {object} model.User
// @Router /api/v1/file/upload [post]
func (h *UserHandler) UploadFile(r *ginext.Request) (*ginext.Response, error) {
	file, err := r.GinCtx.FormFile("data")
	if err != nil {
		return nil, err
	}

	src, err := file.Open()
	if err != nil {
		panic(err)
	}

	// check content type
	buff := make([]byte, 512)
	_, err = src.Read(buff)

	if err != nil {
		panic(err)
	}

	filetype := http.DetectContentType(buff)
	if !utils.IsImageType(filetype) {
		r.GinCtx.JSON(http.StatusBadRequest, "Content type must be an image!")
		return nil, nil
	}

	dst, err := os.Create(filepath.Join("tmp", filepath.Base(file.Filename))) // dir is directory where you want to save file.
	if err != nil {
		panic(err)
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		panic(err)
	}

	//err = h.service.UploadFile(r.GinCtx, model.RegisterRequest{})
	//if err != nil {
	//	return nil, err
	//}

	return ginext.NewResponseData(http.StatusOK, nil), err
}
