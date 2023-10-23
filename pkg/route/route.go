package route

import (
	"finan/mvt-adapter/pkg/handler"
	"finan/mvt-adapter/pkg/repo"
	service2 "finan/mvt-adapter/pkg/service"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"gitlab.com/goxp/cloud0/ginext"
	"gitlab.com/goxp/cloud0/service"
)

type Service struct {
	*service.BaseApp
}

func NewService() *Service {
	s := &Service{
		service.NewApp("Service user", "v1.0"),
	}

	validate := validator.New()

	// repo
	db := s.GetDB()
	repoPG := repo.NewPGRepo(db)

	// service
	userService := service2.NewUserService(repoPG)

	// handle
	handlers := handler.NewUserHandler(userService, validate)

	v1Api := s.Router.Group("/api/v1")
	swaggerApi := s.Router.Group("/")

	// swagger
	swaggerApi.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	v1Api.POST("/user/register", ginext.WrapHandler(handlers.Register))
	v1Api.POST("/user/login", ginext.WrapHandler(handlers.Login))

	// file
	v1Api.POST("/file/upload", ginext.WrapHandler(handlers.Login))

	// Migrate
	return s
}
