package handler

import (
	"finan/mvt-adapter/pkg/model"
	"finan/mvt-adapter/pkg/service"
	"github.com/go-playground/validator/v10"
	"gitlab.com/goxp/cloud0/ginext"
	"net/http"
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
	r.MustBind(&req)

	rs, err := h.service.Login(r.GinCtx, req)
	if err != nil {
		return nil, err
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
	r.MustBind(&req)

	rs, err := h.service.Register(r.GinCtx, req)
	if err != nil {
		return nil, err
	}

	return ginext.NewResponseData(http.StatusOK, rs), nil
}
